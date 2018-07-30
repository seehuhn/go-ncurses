package ncurses

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// #include <ncurses.h>
import "C"

// SetTimeout sets blocking or non-blocking read for a given window.
// If `delay` is negative, blocking read is used (i.e., waits
// indefinitely for input).  If `delay` is zero, then non-blocking
// read is used (i.e., `GetCh()` returns `KeyTimeout` if no input is
// waiting).  If `delay` is positive, then read blocks for `delay`
// milliseconds, and returns `KeyTimeout` if there is still no input.
func (w *Window) SetTimeout(delay int) {
	C.wtimeout(w.ptr, C.int(delay))
	w.timeout = delay
}

// GetTimeout returns the currently used timeout for input.  See
// `SetTimeout()` for the meaning of the returned value.
func (w *Window) GetTimeout() int {
	return w.timeout
}

// GetCh reads a character from the window.  Input is expected to be
// utf-8 encoded and is converted to `rune`.
//
// Function key presses (e.g. cursor keys) are reported as runes in
// the "unicode private use area".  The constants KeyA1, ..., KeyF24
// give all supported function keys.
//
// Refresh() is called before any character is read.
func (w *Window) GetCh() rune {
	var buf []byte
	for !utf8.FullRune(buf) {
		c := C.wgetch(w.ptr)

		if c == C.ERR {
			return KeyTimeout
		}

		if len(buf) == 0 {
			key, found := keyLookup[c]
			if found {
				return key
			}
		}

		if c < 0 || c > 255 {
			return 0
		}
		buf = append(buf, byte(c))
	}
	res, _ := utf8.DecodeRune(buf)
	return res
}

func control(c rune) rune {
	return c - 'A' + 1
}

// Readline allows the user to enter a line of text.  Simple
// line-editing is provided and up to `maxLen` (unicode) characters of
// text can be entered.
//
// Refresh() is called before any character is read.
func (w *Window) Readline(maxLen int) string {
	oldCurs, cerr := CursSet(CursorOn)

	y, x := w.GetYX()
	_, width := w.GetMaxYX()
	if x+maxLen > width {
		maxLen = width - x
	}

	var res []rune
	pos := 0
	for {
		s := string(res)
		s = fmt.Sprintf("%-*s", maxLen, s)
		w.MvAddStr(y, x, s)
		if pos >= maxLen {
			w.Move(y, x+maxLen-1)
		} else {
			w.Move(y, x+pos)
		}
		c := w.GetCh()

		if unicode.IsPrint(c) {
			if len(res) < maxLen {
				res = append(res[:pos], append([]rune{c}, res[pos:]...)...)
				pos++
			} else {
				Beep()
			}
		} else if c == KeyEnter || c == '\n' || c == '\r' {
			break
		} else if c == KeyLeft {
			if pos > 0 {
				pos--
			} else {
				Beep()
			}
		} else if c == KeyRight {
			if pos < len(res) {
				pos++
			} else {
				Beep()
			}
		} else if c == KeyDc || c == control('D') {
			if pos < len(res) {
				res = append(res[:pos], res[pos+1:]...)
			} else {
				Beep()
			}
		} else if c == KeyBackspace || c == 127 || c == control('H') {
			if pos > 0 {
				res = append(res[:pos-1], res[pos:]...)
				pos--
			} else {
				Beep()
			}
		} else if c == control('A') {
			pos = 0
		} else if c == control('E') {
			pos = len(res)
		} else if c == control('K') {
			res = res[:pos]
		} else if c == control('T') {
			if pos > 0 && len(res) > 1 {
				k := pos
				if k > len(res)-1 {
					k = len(res) - 1
				}
				res[k-1], res[k] = res[k], res[k-1]
			}
		} else if c == 27 {
			delay := w.GetTimeout()
			w.SetTimeout(0)
			c2 := w.GetCh()
			w.SetTimeout(delay)

			switch c2 {
			case KeyBackspace, 127, control('H'):
				old := pos
				skip := true
				for pos > 0 &&
					(skip ||
						unicode.IsLetter(res[pos-1]) ||
						unicode.IsDigit(res[pos-1])) {
					skip = skip &&
						!unicode.IsLetter(res[pos-1]) &&
						!unicode.IsDigit(res[pos-1])
					pos--
				}
				res = append(res[:pos], res[old:]...)
			case 'd':
				old := pos
				skip := true
				for pos < len(res) &&
					(skip ||
						unicode.IsLetter(res[pos]) ||
						unicode.IsDigit(res[pos])) {
					skip = skip &&
						!unicode.IsLetter(res[pos]) &&
						!unicode.IsDigit(res[pos])
					pos++
				}
				res = append(res[:old], res[pos:]...)
				pos = old
			default:
				Beep()
			}
		} else {
			Beep()
		}
	}

	if cerr == nil {
		CursSet(oldCurs)
	}

	return string(res)
}
