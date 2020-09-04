// seehuhn.de/go/ncurses - a Go-wrapper for the ncurses library
// Copyright (C) 2018  Jochen Voss <voss@seehuhn.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ncurses

// #cgo CFLAGS: -D_XOPEN_SOURCE -D_XOPEN_SOURCE_EXTENDED
// #cgo LDFLAGS: -lncursesw
//
// #include <locale.h>
// #include <ncurses.h>
//
// WINDOW *
// term_init()
// {
//     WINDOW *scr;
//     setlocale(LC_ALL, "");
//     scr = initscr();
//     start_color();
//
//     cbreak();
//     noecho();
//     keypad(scr, TRUE);
//
//     return scr;
// }
import "C"

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

var global struct {
	sync.Mutex

	isInitialized bool
	signals       chan os.Signal
}

// Beep is used to alert the terminal user.  The function sounds an
// audible alarm on the terminal, if possible; otherwise it flashes
// the screen.
func Beep() {
	C.beep()
}

// EndWin must be called before the program exits, in order to restore
// the terminal to a usable state.
func EndWin() {
	signal.Stop(global.signals)
	close(global.signals)

	C.endwin()

	global.Lock()
	if !global.isInitialized {
		panic("ncurses not initialized")
	}
	global.isInitialized = false
	global.Unlock()
}

// Init initialises the curses library and returns a Window
// corresponding to the whole screen.  You can either use this window,
// or allocate your own, smaller windows using NewWin().
func Init() *Window {
	global.Lock()
	defer global.Unlock()

	if global.isInitialized {
		panic("ncurses already initialized")
	}
	global.isInitialized = true

	// The signal handlers must be installed before term_init() is called.
	global.signals = make(chan os.Signal, 1)
	signal.Notify(global.signals,
		syscall.SIGCONT,
		syscall.SIGTSTP,
		syscall.SIGWINCH)

	res := &Window{
		ptr:     C.term_init(),
		timeout: -1,
	}

	go signalHandler()

	return res
}

func signalHandler() {
	for sig := range global.signals {
		switch sig {
		case syscall.SIGWINCH:
			ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
			if err == nil {
				C.resizeterm(C.int(ws.Row), C.int(ws.Col))
			}
		case syscall.SIGTSTP:
			C.endwin()
		case syscall.SIGCONT:
			C.doupdate()
		}
	}
}

// A Window is the central data structure in the ncurses library.
// Most functionality is implemented as methods of Window objects.
type Window struct {
	ptr     *C.WINDOW
	timeout int
}

// NewWin creates a new window at screen position (beginY, beginX).
// The newly created window has `nLines' lines and `nCols' columns.
// Do not use overlapping windows.
func NewWin(nLines, nCols, beginY, beginX int) *Window {
	nl := C.int(nLines)
	nc := C.int(nCols)
	by := C.int(beginY)
	bx := C.int(beginX)
	ptr := C.newwin(nl, nc, by, bx)
	if ptr == nil {
		return nil
	}
	return &Window{
		ptr:     ptr,
		timeout: -1,
	}
}

// ScrollOk controls what happens when the cursor of a window is moved
// off the edge of the window or scrolling region, either as a result
// of a newline action on the bottom line, or typing the last
// character of the last line.  If disabled, (bf is `false`), the cursor
// is left on the bottom line.  If enabled, (bf is `true`), the window
// is scrolled up one line (Note that to get the physical scrolling
// effect on the terminal, it is also necessary to call IdlOk).
func (w *Window) ScrollOk(bf bool) {
	C.scrollok(w.ptr, C.bool(bf))
}

// IdlOk can be used to allow curses to use the insert/delete line
// feature of terminals so equipped.  If IdlOk is called with `true`
// as second argument, curses considers using insert/delete line.
// Calling idlok with `false` as second argument disables use of line
// insertion and deletion.  This option should be enabled only if the
// application needs insert/delete line, for example, for a screen
// editor.  It is disabled by default because insert/delete line tends
// to be visually annoying when used in applications where it is not
// really needed.  If insert/delete line cannot be used, curses
// redraws the changed portions of all lines.
func (w *Window) IdlOk(bf bool) {
	C.idlok(w.ptr, C.bool(bf))
}

// GetYX returns the current cursor position in the given window.  The
// returned values are the current row y and column x, relative to the
// top-left corner of the window.
func (w *Window) GetYX() (int, int) {
	x := C.getcurx(w.ptr)
	y := C.getcury(w.ptr)
	return int(y), int(x)
}

// GetBegYX returns the coordinates of the top-left corner of the
// window in screen coordinates.  The returned values are the current
// row y and column x, relative to the top-left corner of the screen.
func (w *Window) GetBegYX() (int, int) {
	x := C.getbegx(w.ptr)
	y := C.getbegy(w.ptr)
	return int(y), int(x)
}

// GetMaxYX returns the width and height of the window in characters.
func (w *Window) GetMaxYX() (int, int) {
	x := C.getmaxx(w.ptr)
	y := C.getmaxy(w.ptr)
	return int(y), int(x)
}

// SetBackground manipulates the background of the named window.  The
// window background consists of a combination of attributes (i.e.,
// rendition) and a complex character.  The attribute part of the
// background is combined (or'ed) with all non-blank characters that
// are written into the window.  Both the character and attribute
// parts of the background are combined with the blank characters.
// The background becomes a property of the character and moves with
// the character through any scrolling and insert/delete
// line/character operations.
//
// To the extent possible on a particular terminal, the attribute part
// of the background is displayed as the graphic rendition of the
// character put on the screen.
func (w *Window) SetBackground(char string, attrs AttrType, colorPair ColorPair) {
	var c C.cchar_t
	wch := stringToC(char)
	C.setcchar(&c, &wch[0], C.attr_t(attrs), C.short(colorPair), C.NULL)
	C.wbkgrndset(w.ptr, &c)
}
