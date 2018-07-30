// attrs.go - screen attribute handling for the ncurses package
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
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package ncurses

// #include <ncurses.h>
import "C"

// AttrType describes display attributes (like bold face or blinking)
// for text printed on the screen.
type AttrType int

// These constants list the display attributes supported by ncurses.
// Attributes can be combined by taking the bitwise or of individual
// attributes.
var (
	AttrNormal     AttrType = C.A_NORMAL     // Normal display (no highlight)
	AttrStandout   AttrType = C.A_STANDOUT   // Best highlighting mode of the terminal.
	AttrUnderline  AttrType = C.A_UNDERLINE  // Underlining
	AttrReverse    AttrType = C.A_REVERSE    // Reverse video
	AttrBlink      AttrType = C.A_BLINK      // Blinking
	AttrDim        AttrType = C.A_DIM        // Half bright
	AttrBold       AttrType = C.A_BOLD       // Extra bright or bold
	AttrProtect    AttrType = C.A_PROTECT    // Protected mode
	AttrInvis      AttrType = C.A_INVIS      // Invisible or blank mode
	AttrAltcharset AttrType = C.A_ALTCHARSET // Alternate character set
)

// AttrSet sets the current attributes of the given window to `attrs`.
// All attributes can be turned off using AttrSet(AttrNormal).
func (w *Window) AttrSet(attrs AttrType) {
	C.wattrset(w.ptr, C.int(attrs))
}

// AttrGet returns the current display attributes for the given
// window.
func (w *Window) AttrGet() AttrType {
	var attrs C.attr_t
	C.wattr_get(w.ptr, &attrs, nil, nil)
	return AttrType(attrs)
}

// AttrOn turns on the named attributes without affecting any others.
func (w *Window) AttrOn(attrs AttrType) {
	C.wattron(w.ptr, C.int(attrs))
}

// AttrOff turns off the named attributes without turning any other
// attributes on or off.
func (w *Window) AttrOff(attrs AttrType) {
	C.wattroff(w.ptr, C.int(attrs))
}
