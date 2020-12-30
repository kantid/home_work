package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const doubleslash = '\\'

var ErrInvalidString = errors.New("invalid string")
var ErrorOutOfRange = errors.New("out of range")

func Unpack(s string) (string, error) {
	var inc = []rune(s)
	var whr strings.Builder

	if len(inc) == 0 {
		return "", nil
	}

	for i := 0; i < len(inc); i++ {
		curRune := inc[i]
		count := 1

		if curRune == doubleslash {
			if escaped, err := ReturnFromSlice(inc, i); err != nil {
				return "", ErrInvalidString
			} else if !escaped {
				continue
			}
		}

		if unicode.IsDigit(curRune) {
			if escaped, err := ReturnFromSlice(inc, i); err != nil || !escaped {
				return "", ErrInvalidString
			}
		}

		if i+1 < len(inc) && unicode.IsDigit(inc[i+1]) {
			var err error
			if count, err = strconv.Atoi(string(inc[i+1])); err != nil {
				return "", ErrInvalidString
			}
			i++
		}

		whr.WriteString(strings.Repeat(string(curRune), count))
	}

	return whr.String(), nil
}

func ReturnFromSlice(j []rune, i int) (bool, error) {
	if i == 0 {
		return false, nil
	}
	if i < 0 || i >= len(j) {
		return false, ErrorOutOfRange
	}
	if j[i-1] != doubleslash {
		return false, nil
	}

	i--
	backSlashEscaped, _ := ReturnFromSlice(j, i)
	return !backSlashEscaped, nil
}
