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
	"github.com/seehuhn/vocab/ncurses"
)

func main() {
	win := ncurses.Init()
	defer ncurses.EndWin()

	win.SetTimeout(1000)

	for {
		win.AddStr("\npress any key (q to exit)\n")
		c := win.GetCh()
		if c == 'q' {
			break
		} else if c == ncurses.KeyResize {
			win.Println("resize")
		} else if c == ncurses.KeyTimeout {
			delay := win.GetTimeout()
			win.Println("*timeout", delay, "*")
			win.SetTimeout(2 * delay)
			continue
		} else if c > 0 {
			win.Printf("%c %d\n", c, c)
		} else {
			win.Printf("* %d\n", c)
		}
	}
}
