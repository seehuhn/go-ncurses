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

// #include <ncurses.h>
import "C"

// CursorVisibility describes the state of the cursor.
type CursorVisibility int

// These constants give the possible values for CursorVisibility.
const (
	CursorOff CursorVisibility = iota
	CursorOn
	CursorStrong
)

// CursSet sets the cursor state to invisible, normal, or very
// visible, depending on the value of `visibility`.  If the terminal
// supports the visibility requested, the previous cursor state is
// returned; otherwise an error is returned.
func CursSet(visibility CursorVisibility) (CursorVisibility, error) {
	res := C.curs_set(C.int(visibility))
	if res == C.ERR {
		return CursorOff, ErrNotSupported
	}
	return CursorVisibility(res), nil
}
