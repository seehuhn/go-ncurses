package ncurses

import "errors"

// Errors returned by the ncurses routines.
var (
	ErrNotSupported = errors.New("setting not supported")
	ErrColorFailed  = errors.New("color setting failed")
)
