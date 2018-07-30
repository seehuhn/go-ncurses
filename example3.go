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
