// +build ignore

package main

import (
	"fmt"

	"github.com/seehuhn/vocab/ncurses"
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
