package terminal

// ANSI color values
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGrey
	DarkGrey
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

// Default colors are potentially distinct to allow for special behavior.
// For example, a transparent background. Otherwise, the simple case is to
// map default colors to another color.
const (
	DefaultFG Color = DefaultColor
	DefaultBG Color = DefaultColor
)

// Color maps to the ANSI colors [0, 16) and the xterm colors [16, 256). If IsRGBColor is set, then the first 24 bytes are the R, G, and B components.
type Color uint32

// Flags to signify if the color is actually an RGB value or a default color.
const (
	RGBColor Color = 1 << 24
	DefaultColor Color = 1 << 25
)

// ANSI returns true if Color is within [0, 16).
func (c Color) ANSI() bool {
	return (c < 16)
}

// RGB returns the stored red, green, and blue components. Only valid if RGBColor is set.
func (c Color) RGB() (uint8, uint8, uint8) {
	return uint8((c >> 16) & 0xff), uint8((c >> 8) & 0xff), uint8(c & 0xff)
}
