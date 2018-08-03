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

import "seehuhn.de/go/ncurses"

func main() {
	ncurses.Init()
	defer ncurses.EndWin()

	c1 := ncurses.ColorPair(1)
	c1.Init(ncurses.ColorRed, ncurses.ColorYellow)
	c2 := ncurses.ColorPair(2)
	c2.Init(ncurses.ColorGreen, ncurses.ColorBlue)

	w1 := ncurses.NewWin(5, 10, 1, 0)
	w1.SetBackground("X", ncurses.AttrBold, c1)
	w1.Erase()
	w1.Println("some text")
	w1.Refresh()

	w2 := ncurses.NewWin(5, 10, 0, 10)
	w2.SetBackground(".", ncurses.AttrBlink, c2)
	w2.Erase()

	w2.AttrSet(c1.AsAttr())
	w2.Println("some text")
	w2.Refresh()

	w3 := ncurses.NewWin(1, 20, 6, 0)
	w3.Print("press any key ...")
	w3.GetCh()
}
