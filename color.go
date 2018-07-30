package ncurses

// #include <ncurses.h>
import "C"

// Color describes the text and background colors supported by
// ncurses.
type Color int

// These constants define default colors, corresponding to color
// values 0, 1, ..., 7.  Additional color values in the range 8, ...,
// NumColors()-1 can be used after initializing them with
// Color.Init().
var (
	ColorBlack   Color = C.COLOR_BLACK
	ColorRed     Color = C.COLOR_RED
	ColorGreen   Color = C.COLOR_GREEN
	ColorYellow  Color = C.COLOR_YELLOW
	ColorBlue    Color = C.COLOR_BLUE
	ColorMagenta Color = C.COLOR_MAGENTA
	ColorCyan    Color = C.COLOR_CYAN
	ColorWhite   Color = C.COLOR_WHITE
)

// NumColors returns the maximum number of colors the terminal can
// support.
func NumColors() int {
	return int(C.COLORS)
}

// Init changes the definition of a Color.  The value color must be in
// the range from 0 to NumColors()-1.  The three arguments are RGB
// values (for the amounts of red, green, and blue components) in the
// range from 0 to 1000.  When Init() is used, all occurrences of that
// color on the screen immediately change to the new definition.
func (color Color) Init(red, green, blue int) error {
	rc := C.init_color(C.short(color),
		C.short(red), C.short(green), C.short(blue))
	if rc == C.ERR {
		return ErrColorFailed
	}
	return nil
}

// NumColorPairs returns the maximum number of color-pairs the
// terminal can support.
func NumColorPairs() int {
	return int(C.COLOR_PAIRS)
}

// A ColorPair represents a combination of foreground and background
// color.  The default color-pair corresponds to the value 0.
// Additional color-pairs in the range 1, ..., NumColorPairs()-1 can
// be initialized using ColorPair.Init().
type ColorPair int

// Init changes the definition of a color-pair.  The arguments give
// the new foreground color and background color.  If the color-pair
// was previously initialized, the screen is refreshed and all
// occurrences of that color-pair are changed to the new definition.
func (pair ColorPair) Init(fg, bg Color) error {
	rc := C.init_pair(C.short(pair), C.short(fg), C.short(bg))
	if rc == C.ERR {
		return ErrColorFailed
	}
	return nil
}

// AsAttr returns a new video attribute, corresponding to the color
// pair.
func (pair ColorPair) AsAttr() AttrType {
	return AttrType(C.COLOR_PAIR(C.int(pair)))
}
