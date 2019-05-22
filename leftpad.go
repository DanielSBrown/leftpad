package leftpad

import (
	"strings"
	"unicode/utf8"
)

var default_char = ' '

// Format takes in a string and an int, and returns string left padded with spaces to the length of the int. If string is already longer than length, the original string is returned
func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

// FormatRune takes in a string, int and a rune and returns the string left-padded with the specific rune to the length of hte int. If string is longer, orig is retured
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	return strings.Repeat(string(r), size-l) + s
}
