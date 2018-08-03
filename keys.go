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

// #include <ncurses.h>
import "C"

// Codes in the "unicode private use area" denote function keys.
const (
	KeyRangeStart rune = '\uE000'

	KeyA1        = KeyRangeStart + iota // upper left of keypad
	KeyA3                               // upper right of keypad
	KeyB2                               // center of keypad
	KeyBackspace                        // backspace key
	KeyBeg                              // begin key
	KeyBreak                            // Break key (unreliable)
	KeyBtab                             // back-tab key
	KeyC1                               // lower left of keypad
	KeyC3                               // lower right of keypad
	KeyCancel                           // cancel key
	KeyCatab                            // clear-all-tabs key
	KeyClear                            // clear-screen or erase key
	KeyClose                            // close key
	KeyCommand                          // command key
	KeyCopy                             // copy key
	KeyCreate                           // create key
	KeyCtab                             // clear-tab key
	KeyDc                               // delete-character key
	KeyDl                               // delete-line key
	KeyDown                             // down-arrow key
	KeyEic                              // sent by rmir or smir in insert mode
	KeyEnd                              // end key
	KeyEnter                            // enter/send key
	KeyEol                              // clear-to-end-of-line key
	KeyEos                              // clear-to-end-of-screen key
	KeyEvent                            // We were interrupted by an event
	KeyExit                             // exit key
	KeyFind                             // find key
	KeyHelp                             // help key
	KeyHome                             // home key
	KeyIc                               // insert-character key
	KeyIl                               // insert-line key
	KeyLeft                             // left-arrow key
	KeyLl                               // lower-left key (home down)
	KeyMark                             // mark key
	KeyMessage                          // message key
	KeyMouse                            // Mouse event has occurred
	KeyMove                             // move key
	KeyNext                             // next key
	KeyNpage                            // next-page key
	KeyOpen                             // open key
	KeyOptions                          // options key
	KeyPpage                            // previous-page key
	KeyPrevious                         // previous key
	KeyPrint                            // print key
	KeyRedo                             // redo key
	KeyReference                        // reference key
	KeyRefresh                          // refresh key
	KeyReplace                          // replace key
	KeyReset                            // Reset or hard reset (unreliable)
	KeyResize                           // Terminal resize event
	KeyRestart                          // restart key
	KeyResume                           // resume key
	KeyRight                            // right-arrow key
	KeySave                             // save key
	KeySelect                           // select key
	KeySf                               // scroll-forward key
	KeySr                               // scroll-backward key
	KeySreset                           // Soft (partial) reset (unreliable)
	KeyStab                             // set-tab key
	KeySuspend                          // suspend key
	KeyUndo                             // undo key
	KeyUp                               // up-arrow key
	KeySBeg                             // shifted begin key
	KeySCancel                          // shifted cancel key
	KeySCommand                         // shifted command key
	KeySCopy                            // shifted copy key
	KeySCreate                          // shifted create key
	KeySDc                              // shifted delete-character key
	KeySDl                              // shifted delete-line key
	KeySEnd                             // shifted end key
	KeySEol                             // shifted clear-to-end-of-line key
	KeySExit                            // shifted exit key
	KeySFind                            // shifted find key
	KeySHelp                            // shifted help key
	KeySHome                            // shifted home key
	KeySIc                              // shifted insert-character key
	KeySLeft                            // shifted left-arrow key
	KeySMessage                         // shifted message key
	KeySMove                            // shifted move key
	KeySNext                            // shifted next key
	KeySOptions                         // shifted options key
	KeySPrevious                        // shifted previous key
	KeySPrint                           // shifted print key
	KeySRedo                            // shifted redo key
	KeySReplace                         // shifted replace key
	KeySRight                           // shifted right-arrow key
	KeySRsume                           // shifted resume key
	KeySSave                            // shifted save key
	KeySSuspend                         // shifted suspend key
	KeySUndo                            // shifted undo key
	KeyTimeout                          // (no input available at timeout)
	KeyF1                               // function key F1
	KeyF2                               // function key F2
	KeyF3                               // function key F3
	KeyF4                               // function key F4
	KeyF5                               // function key F5
	KeyF6                               // function key F6
	KeyF7                               // function key F7
	KeyF8                               // function key F8
	KeyF9                               // function key F9
	KeyF10                              // function key F10
	KeyF11                              // function key F11
	KeyF12                              // function key F12
	KeyF13                              // function key F13
	KeyF14                              // function key F14
	KeyF15                              // function key F15
	KeyF16                              // function key F16
	KeyF17                              // function key F17
	KeyF18                              // function key F18
	KeyF19                              // function key F19
	KeyF20                              // function key F20
	KeyF21                              // function key F21
	KeyF22                              // function key F22
	KeyF23                              // function key F23
	KeyF24                              // function key F24
	KeyRangeEnd
)

var keyLookup map[C.int]rune = map[C.int]rune{
	C.KEY_A1:        KeyA1,
	C.KEY_A3:        KeyA3,
	C.KEY_B2:        KeyB2,
	C.KEY_BACKSPACE: KeyBackspace,
	C.KEY_BEG:       KeyBeg,
	C.KEY_BREAK:     KeyBreak,
	C.KEY_BTAB:      KeyBtab,
	C.KEY_C1:        KeyC1,
	C.KEY_C3:        KeyC3,
	C.KEY_CANCEL:    KeyCancel,
	C.KEY_CATAB:     KeyCatab,
	C.KEY_CLEAR:     KeyClear,
	C.KEY_CLOSE:     KeyClose,
	C.KEY_COMMAND:   KeyCommand,
	C.KEY_COPY:      KeyCopy,
	C.KEY_CREATE:    KeyCreate,
	C.KEY_CTAB:      KeyCtab,
	C.KEY_DC:        KeyDc,
	C.KEY_DL:        KeyDl,
	C.KEY_DOWN:      KeyDown,
	C.KEY_EIC:       KeyEic,
	C.KEY_END:       KeyEnd,
	C.KEY_ENTER:     KeyEnter,
	C.KEY_EOL:       KeyEol,
	C.KEY_EOS:       KeyEos,
	C.KEY_EVENT:     KeyEvent,
	C.KEY_EXIT:      KeyExit,
	C.KEY_F0 + 10:   KeyF10,
	C.KEY_F0 + 11:   KeyF11,
	C.KEY_F0 + 12:   KeyF12,
	C.KEY_F0 + 13:   KeyF13,
	C.KEY_F0 + 14:   KeyF14,
	C.KEY_F0 + 15:   KeyF15,
	C.KEY_F0 + 16:   KeyF16,
	C.KEY_F0 + 17:   KeyF17,
	C.KEY_F0 + 18:   KeyF18,
	C.KEY_F0 + 19:   KeyF19,
	C.KEY_F0 + 1:    KeyF1,
	C.KEY_F0 + 20:   KeyF20,
	C.KEY_F0 + 21:   KeyF21,
	C.KEY_F0 + 22:   KeyF22,
	C.KEY_F0 + 23:   KeyF23,
	C.KEY_F0 + 24:   KeyF24,
	C.KEY_F0 + 2:    KeyF2,
	C.KEY_F0 + 3:    KeyF3,
	C.KEY_F0 + 4:    KeyF4,
	C.KEY_F0 + 5:    KeyF5,
	C.KEY_F0 + 6:    KeyF6,
	C.KEY_F0 + 7:    KeyF7,
	C.KEY_F0 + 8:    KeyF8,
	C.KEY_F0 + 9:    KeyF9,
	C.KEY_FIND:      KeyFind,
	C.KEY_HELP:      KeyHelp,
	C.KEY_HOME:      KeyHome,
	C.KEY_IC:        KeyIc,
	C.KEY_IL:        KeyIl,
	C.KEY_LEFT:      KeyLeft,
	C.KEY_LL:        KeyLl,
	C.KEY_MARK:      KeyMark,
	C.KEY_MESSAGE:   KeyMessage,
	C.KEY_MOUSE:     KeyMouse,
	C.KEY_MOVE:      KeyMove,
	C.KEY_NEXT:      KeyNext,
	C.KEY_NPAGE:     KeyNpage,
	C.KEY_OPEN:      KeyOpen,
	C.KEY_OPTIONS:   KeyOptions,
	C.KEY_PPAGE:     KeyPpage,
	C.KEY_PREVIOUS:  KeyPrevious,
	C.KEY_PRINT:     KeyPrint,
	C.KEY_REDO:      KeyRedo,
	C.KEY_REFERENCE: KeyReference,
	C.KEY_REFRESH:   KeyRefresh,
	C.KEY_REPLACE:   KeyReplace,
	C.KEY_RESET:     KeyReset,
	C.KEY_RESIZE:    KeyResize,
	C.KEY_RESTART:   KeyRestart,
	C.KEY_RESUME:    KeyResume,
	C.KEY_RIGHT:     KeyRight,
	C.KEY_SAVE:      KeySave,
	C.KEY_SBEG:      KeySBeg,
	C.KEY_SCANCEL:   KeySCancel,
	C.KEY_SCOMMAND:  KeySCommand,
	C.KEY_SCOPY:     KeySCopy,
	C.KEY_SCREATE:   KeySCreate,
	C.KEY_SDC:       KeySDc,
	C.KEY_SDL:       KeySDl,
	C.KEY_SELECT:    KeySelect,
	C.KEY_SEND:      KeySEnd,
	C.KEY_SEOL:      KeySEol,
	C.KEY_SEXIT:     KeySExit,
	C.KEY_SF:        KeySf,
	C.KEY_SFIND:     KeySFind,
	C.KEY_SHELP:     KeySHelp,
	C.KEY_SHOME:     KeySHome,
	C.KEY_SIC:       KeySIc,
	C.KEY_SLEFT:     KeySLeft,
	C.KEY_SMESSAGE:  KeySMessage,
	C.KEY_SMOVE:     KeySMove,
	C.KEY_SNEXT:     KeySNext,
	C.KEY_SOPTIONS:  KeySOptions,
	C.KEY_SPREVIOUS: KeySPrevious,
	C.KEY_SPRINT:    KeySPrint,
	C.KEY_SR:        KeySr,
	C.KEY_SREDO:     KeySRedo,
	C.KEY_SREPLACE:  KeySReplace,
	C.KEY_SRESET:    KeySreset,
	C.KEY_SRIGHT:    KeySRight,
	C.KEY_SRSUME:    KeySRsume,
	C.KEY_SSAVE:     KeySSave,
	C.KEY_SSUSPEND:  KeySSuspend,
	C.KEY_STAB:      KeyStab,
	C.KEY_SUNDO:     KeySUndo,
	C.KEY_SUSPEND:   KeySuspend,
	C.KEY_UNDO:      KeyUndo,
	C.KEY_UP:        KeyUp,
}

// #define KEY_F(n)	(KEY_F0+(n))	// Value of function key n
