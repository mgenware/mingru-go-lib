package mingru

import "strings"

// InputPlaceholders returns a number of SQL input placeholders (?) based on the given count.
func InputPlaceholders(count int) string {
	if count <= 1 {
		return "?"
	}
	return "?" + strings.Repeat(", ?", count-1)
}
