// Package ncurses is a unicode-aware wrapper for the ncurses library.
//
// If only a single window is required, the library can be initialized
// as follows:
//
//     win := ncurses.Init()
//     defer ncurses.EndWin()
//     // use methods of win for input and output
//
// If more than one window is required, the window returned by
// term.Init() can be ignored and (non-overlapping) windows can be
// allocated using term.NewWin() instead:
//
//     win := ncurses.Init()
//     defer ncurses.EndWin()
//     height, width := win.GetMaxYX()
//     win1 := ncurses.NewWin(5, width, 0, 0)
//     win2 := ncurses.NewWin(height-5, width, 5, 0)
//     // Use methods of win1 and win2 for input and output.
//     // Do NOT use win after this point.
//
// Screen coordinates follow the ncurses conventions: in argument
// lists, the row y is given first, followed by the column x.  The top
// left corner of a window coresponds to y=0, x=0.
package ncurses
