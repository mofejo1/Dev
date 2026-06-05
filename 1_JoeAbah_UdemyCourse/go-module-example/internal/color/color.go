package color

import (
	"fmt"
	"strconv"
	"strings"
)

type Color struct {
	value int
}

var (
	Black   = Color{value: 30}
	Red     = Color{value: 31}
	Green   = Color{value: 32}
	Yellow  = Color{value: 33}
	Blue    = Color{value: 34}
	Magenta = Color{value: 35}
	Cyan    = Color{value: 36}
	White   = Color{value: 37}

	Bold      = Color{value: 1}
	underline = Color{value: 4}
)

func Text(text string, color ...Color) string {

	if len(color) == 0 {
		return text
	}
	var codes []string

	for _, color := range color {
		// codes = append(codes, fmt.Sprintf("%d", color.value))
		codes = append(codes, strconv.Itoa(color.value))

	}
	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(codes, ";"), text)
}
