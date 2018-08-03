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

// Package ncurses is a unicode-aware wrapper for the ncurses library.
//
// If only a single window is required, the library can be initialized
// as follows:
//
//     win := ncurses.Init()
//     defer ncurses.EndWin()
//     // use methods of win for input and output
//
// If more than one window is required, the window returned by
// term.Init() can be ignored and (non-overlapping) windows can be
// allocated using term.NewWin() instead:
//
//     win := ncurses.Init()
//     defer ncurses.EndWin()
//     height, width := win.GetMaxYX()
//     win1 := ncurses.NewWin(5, width, 0, 0)
//     win2 := ncurses.NewWin(height-5, width, 5, 0)
//     // Use methods of win1 and win2 for input and output.
//     // Do NOT use win after this point.
//
// Screen coordinates follow the ncurses conventions: in argument
// lists, the row y is given first, followed by the column x.  The top
// left corner of a window coresponds to y=0, x=0.
package ncurses // import "seehuhn.de/go/ncurses"
