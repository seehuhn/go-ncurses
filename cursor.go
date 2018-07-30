package ncurses

// #include <ncurses.h>
import "C"

// CursorVisibility describes the state of the cursor.
type CursorVisibility int

// These constants give the possible values for CursorVisibility.
const (
	CursorOff CursorVisibility = iota
	CursorOn
	CursorStrong
)

// CursSet sets the cursor state to invisible, normal, or very
// visible, depending on the value of `visibility`.  If the terminal
// supports the visibility requested, the previous cursor state is
// returned; otherwise an error is returned.
func CursSet(visibility CursorVisibility) (CursorVisibility, error) {
	res := C.curs_set(C.int(visibility))
	if res == C.ERR {
		return CursorOff, ErrNotSupported
	}
	return CursorVisibility(res), nil
}
