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

// +build ignore

package main

import "github.com/seehuhn/vocab/ncurses"

func main() {
	win := ncurses.Init()
	defer ncurses.EndWin()
	height, width := win.GetMaxYX()
	win1 := ncurses.NewWin(5, width, 0, 0)
	win1.ScrollOk(true)
	win1.AttrSet(ncurses.AttrBold)
	win2 := ncurses.NewWin(height-6, width, 5, 0)
	win2.ScrollOk(true)
	win3 := ncurses.NewWin(1, width, height-1, 0)

	win3.Println("press any key")
	win3.SetTimeout(0)
	for i := 0; ; i++ {
		if i%1000 == 0 {
			win1.Printf("\nline %d ...", i)
			win1.Refresh()
		}
		win2.Println("line", i)
		win2.Refresh()
		c := win3.GetCh()
		if c != ncurses.KeyTimeout {
			break
		}
	}
}
