package component

import (
	"fmt"
	"strings"
)

func PrintComplexSeparator() {
	fmt.Printf("|%s|%s|%s|%s|%s|\n",
		strings.Repeat("=", 22),
		strings.Repeat("=", 22),
		strings.Repeat("=", 32),
		strings.Repeat("=", 12),
		strings.Repeat("=", 12))
}
