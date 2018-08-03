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
	"math"

	"seehuhn.de/go/ncurses"
)

func rgb(alpha float64) (int, int, int) {
	H := alpha * 6
	X := 1 - math.Abs(math.Mod(H, 2)-1)

	var r, g, b float64
	switch {
	case H < 1:
		r, g, b = 1, X, 0
	case H < 2:
		r, g, b = X, 1, 0
	case H < 3:
		r, g, b = 0, 1, X
	case H < 4:
		r, g, b = 0, X, 1
	case H < 5:
		r, g, b = X, 0, 1
	default:
		r, g, b = 1, 0, X
	}
	return int(1000.99 * r), int(1000.99 * g), int(1000.99 * b)
}

func main() {
	win := ncurses.Init()
	defer ncurses.EndWin()

	numColors := ncurses.NumColors()
	win.AddStr(fmt.Sprintf("COLORS = %d\n", numColors))
	numPairs := ncurses.NumColorPairs()
	win.AddStr(fmt.Sprintf("COLOR_PAIRS = %d\n", numPairs))

	if numPairs >= 8 {
		table := []struct {
			Name string
			Col  ncurses.Color
		}{
			{"red", ncurses.ColorRed},
			{"green", ncurses.ColorGreen},
			{"yellow", ncurses.ColorYellow},
			{"blue", ncurses.ColorBlue},
			{"magenta", ncurses.ColorMagenta},
			{"cyan", ncurses.ColorCyan},
			{"white", ncurses.ColorWhite},
		}
		win.Println()
		for i, entry := range table {
			// ColorPair(0) are the default colours
			cp := ncurses.ColorPair(i + 1)
			cp.Init(entry.Col, ncurses.ColorBlack)
			win.AttrOn(cp.AsAttr())
			win.Println(entry.Name)
		}
	}

	k := numColors - 16
	if k > numPairs-8 {
		k = numPairs - 8
	}
	_, width := win.GetMaxYX()
	if k > width-1 {
		k = width - 1
	}
	if k > 1 {
		win.Println()
		for i := 0; i < k; i++ {
			col := ncurses.Color(i + 16)
			r, g, b := rgb(float64(i) / float64(k))
			col.Init(r, g, b)
			cp := ncurses.ColorPair(i + 8)
			cp.Init(col, ncurses.ColorBlack)
		}
		for l := 0; l < 5; l++ {
			for i := 0; i <= k; i++ {
				j := (i + l) % k
				cp := ncurses.ColorPair(j + 8)
				win.AttrSet(cp.AsAttr())
				win.Print("X")
			}
			if k < width-1 {
				win.Println()
			}
		}
	}

	win.AttrSet(0)
	win.AddStr("\npress any key\n")
	win.GetCh()
}
