package bslogger

import (
	"fmt"
)

type ansiEscapeCode int

const (
	// Reset previous codes
	reset ansiEscapeCode = 0

	// Display attributes
	normal           = 22
	bold             = 1
	faint            = 2
	italic           = 3
	noItalic         = 23
	underline        = 4
	doubleUnderline  = 21
	noUnderline      = 24
	invert           = 7
	noInvert         = 27
	strike           = 9
	noStrike         = 29
	fontDefault      = 10
	framed           = 51
	encircled        = 52
	noFrameEncircled = 54

	// Foreground colors
	fgDefault      = 39
	fgBlack        = 30
	fgRed          = 31
	fgGreen        = 32
	fgYellow       = 33
	fgBlue         = 34
	fgPurple       = 35
	fgCyan         = 36
	fgWhite        = 37
	fgBrightBlack  = 90
	fgBrightRed    = 91
	fgBrightGreen  = 92
	fgBrightYellow = 93
	fgBrightBlue   = 94
	fgBrightPurple = 95
	fgBrightCyan   = 96
	fgBrightWhite  = 97

	// Background colors
	bgDefault      = 49
	bgBlack        = 40
	bgRed          = 41
	bgGreen        = 42
	bgYellow       = 43
	bgBlue         = 44
	bgPurple       = 45
	bgCyan         = 46
	bgWhite        = 47
	bgBrightBlack  = 100
	bgBrightRed    = 101
	bgBrightGreen  = 102
	bgBrightYellow = 103
	bgBrightBlue   = 104
	bgBrightPurple = 105
	bgBrightCyan   = 106
	bgBrightWhite  = 107
)

func ansiEscapeEncode(message string, fg ansiEscapeCode, bg ansiEscapeCode, display ansiEscapeCode) string {
	return fmt.Sprintf("\033[%d;%d;%dm%s\033[0m", fg, bg, display, message)
}
