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

import "testing"

func TestAttrs(t *testing.T) {
	win := Init()
	defer EndWin()

	win.AttrSet(AttrBold | AttrBlink)
	win.AttrOn(AttrUnderline | AttrAltcharset)
	win.AttrOff(AttrBlink | AttrAltcharset)
	want := AttrBold | AttrUnderline
	have := win.AttrGet()
	if have != want {
		t.Errorf("wrong attbritutes: expected %08X, got %08X", want, have)
	}
}
