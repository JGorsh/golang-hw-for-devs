package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	var prev string
	var j int

	if unicode.IsDigit(rune(s[0])) {
		return sb.String(), ErrInvalidString
	} else {
		for i, val := range s {
			if ch, err := strconv.Atoi(string(val)); err != nil {
				sb.WriteString(string(val))
				prev = string(s[i])
				j = 0
			} else {
				j++
				if j == 2 {
					return sb.String(), ErrInvalidString
				}
				if ch != 0 {
					sb.WriteString(strings.Repeat(prev, ch-1))
				} else {
					var pr = sb.String()
					pr = pr[:len(pr)-1]
					sb.Reset()
					sb.WriteString(pr)
				}
			}
		}
	}
	return sb.String(), ErrInvalidString
}
