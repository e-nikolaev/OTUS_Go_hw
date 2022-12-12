package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inString string) (string, error) {
	var outString strings.Builder

	inputStringRune := []rune(inString)
	inputStringLen := len(inputStringRune)

	for i := 0; i < inputStringLen; i++ {
		curSymbol := inputStringRune[i]

		if unicode.IsDigit(inputStringRune[i]) {
			return "", ErrInvalidString
		}

		if string(curSymbol) == `\` {
			if i == inputStringLen-1 {
				return "", ErrInvalidString
			}
			i++
			curSymbol = inputStringRune[i]
		}

		// For last symbol in the string
		if i == inputStringLen-1 {
			outString.WriteRune(curSymbol)
			break
		}

		nextSymbol := inputStringRune[i+1]

		if unicode.IsDigit(nextSymbol) {
			count, _ := strconv.Atoi(string(nextSymbol))
			outString.WriteString(strings.Repeat(string(curSymbol), count))
			i++
		} else {
			outString.WriteRune(curSymbol)
		}
	}

	return outString.String(), nil
}
