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
