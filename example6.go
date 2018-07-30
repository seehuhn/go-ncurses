// +build ignore

package main

import "github.com/seehuhn/vocab/ncurses"

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
