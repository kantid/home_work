package hw02_unpack_string

import (
	"errors"
	"strings"
	"unicode"
)

// ErrInvalidString - invalid string error
var ErrInvalidString = errors.New("invalid string")

// Unpack - unpack string
// - s is entire string
// - this func returns string results of Unpack and error, if it's throw
func Unpack(s string) (string, error) {
	// res - results of Unpack function return, we use Builder to min memory usage
	var res strings.Builder
	// lastRune -  last symbol before current
	var lastRune rune
	var esc bool

	//go throw string by every rune, i is current symbol position
	for i, curRune := range s {
		//if first symbol is dijit, error occurred
		if unicode.IsDigit(curRune) && i == 0 {
			return "", ErrInvalidString
		}

		//if we get number instead of digit, error occurred
		if unicode.IsDigit(curRune) && unicode.IsDigit(lastRune) {
			return "", ErrInvalidString
		}

		//if current rune is letter or space character, write current rune to res
		if unicode.IsLetter(curRune) || unicode.IsSpace(curRune) || unicode.IsSymbol(curRune) {
			res.WriteRune(curRune)
		}

		if unicode.IsDigit(curRune) && esc {
			res.WriteRune(curRune)
		}

		//if current symbol is digit & there's letter after, or last symbol = space character
		if (unicode.IsDigit(curRune) && !esc) && (unicode.IsLetter(lastRune) || unicode.IsSpace(lastRune) || unicode.IsSymbol(lastRune)) {
			//convert rune to int
			runeInt := int(curRune - '0')
				//if current rune is '0', remove lastRune from result
				if runeInt == 0 {
					s = strings.TrimSuffix(res.String(), string(lastRune))
					res.Reset()
					res.WriteString(s)
					//if not 0, repeat last rune by current rune value
				} else {
					res.WriteString(strings.Repeat(string(lastRune), runeInt-1))
			}
		}

		if string(curRune) == "\\" {
			esc = true
		} else {
			esc = false
		}

		if esc {}

		// set lastRune to current rune
		lastRune = curRune
	}

	//return result, nil because we don't have error here
	return res.String(), nil

}