// +build ignore

package main

import (
	"fmt"

	"github.com/seehuhn/vocab/ncurses"
)

func main() {
	win := ncurses.Init()
	win.Println("enter any text and press return:")

	n := 40
	win.Move(1, 0)
	win.Print("|")
	win.Move(1, n+1)
	win.Print("|")
	win.Move(1, 1)
	s := win.Readline(n)

	ncurses.EndWin()
	fmt.Println(s)
}
