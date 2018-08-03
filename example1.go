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

import (
	"fmt"

	"seehuhn.de/go/ncurses"
)

func main() {
	win := ncurses.Init()
	defer ncurses.EndWin()

	table := []struct {
		Name string
		Code ncurses.AttrType
	}{
		{"normal", ncurses.AttrNormal},
		{"standout", ncurses.AttrStandout},
		{"underline", ncurses.AttrUnderline},
		{"reverse", ncurses.AttrReverse},
		{"blink", ncurses.AttrBlink},
		{"dim", ncurses.AttrDim},
		{"bold", ncurses.AttrBold},
		{"protect", ncurses.AttrProtect},
		{"invis", ncurses.AttrInvis},
		{"altcharset", ncurses.AttrAltcharset},
	}

	win.AddStr("\n")
	for _, entry := range table {
		win.AddStr(fmt.Sprintf("%12s ", entry.Name))
		win.AttrOn(entry.Code)
		win.AddStr("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
		win.AttrOff(entry.Code)
		win.AddStr("\n")
	}
	win.Refresh()

	win.AddStr("\npress any key\n")
	win.GetCh()
}
