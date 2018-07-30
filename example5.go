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
