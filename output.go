package ncurses

import (
	"fmt"
	"unicode/utf8"
)

// #include <ncurses.h>
import "C"

// Refresh must be called to get actual output to the terminal, as
// other routines merely manipulate data structures.  The routine
// copies the named window to the physical terminal screen, taking
// into account what is already there to do optimizations.
//
// Note that methods waiting for keyboard input call Refresh() before
// waiting for input.
func (w *Window) Refresh() {
	C.wrefresh(w.ptr)
}

// Erase copies blanks to every position in the window, thus clearing
// the window.  The blanks created by Erase() have the current
// background rendition, as set by BkgdSet(), merged into them.
func (w *Window) Erase() {
	C.werase(w.ptr)
}

// ClrToEol erases the current line to the right of the cursor,
// inclusive, to the end of the current line.
func (w *Window) ClrToEol() {
	C.wclrtoeol(w.ptr)
}

// Move changes the current cursor position of the window.  The
// arguments indicate the row (y) and column (x) of the new cursor
// position.
func (w *Window) Move(y, x int) {
	C.wmove(w.ptr, C.int(y), C.int(x))
}

// AddStr prints a string to the window at the current cursor
// position.
func (w *Window) AddStr(s string) {
	buf := stringToC(s)
	C.waddwstr(w.ptr, &buf[0])
}

// MvAddStr moves the cursor to row y, column x, and then prints a
// string to the window at the new cursor position.
func (w *Window) MvAddStr(y, x int, s string) {
	buf := stringToC(s)
	C.mvwaddwstr(w.ptr, C.int(y), C.int(x), &buf[0])
}

// Print formats the arguments using their default formats and writes
// the resulting string to the window.  Spaces are added between
// operands when neither is a string.
func (w *Window) Print(a ...interface{}) {
	w.AddStr(fmt.Sprint(a...))
}

// Printf formats the arguments according to a format specifier and
// writes the resulting string to the window.
func (w *Window) Printf(format string, a ...interface{}) {
	w.AddStr(fmt.Sprintf(format, a...))
}

// Println formats its arguments using their default formats and
// writes the resulting string to the window.  Spaces are always added
// between operands and a newline is appended.
func (w *Window) Println(a ...interface{}) {
	w.AddStr(fmt.Sprintln(a...))
}

func stringToC(s string) []C.wchar_t {
	buf := make([]C.wchar_t, utf8.RuneCountInString(s)+1)
	pos := 0
	for _, c := range s {
		buf[pos] = C.wchar_t(c)
		pos++
	}
	return buf
}
