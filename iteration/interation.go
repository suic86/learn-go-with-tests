package iteration

import "strings"

// Repeat returns the given `character` repeated `count` times. If `count` is less than equal 0 it return an empty string.
func Repeat(character string, count int) string {
	if count <= 0 {
		return ""
	}
	var repeated strings.Builder
	for range count {
		repeated.WriteString(character)
	}
	return repeated.String()
}
