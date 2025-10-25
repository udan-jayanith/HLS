package HLS

import (
	"errors"
	"strings"
)

/*
type AttributeValueType int

const (
	// String of characters from the set
	// [0..9] expressing an integer in base-10 arithmetic in the range
	// from 0 to 2^64-1 (18446744073709551615).  A decimal-integer may be
	// from 1 to 20 characters long.
	DecimalInteger AttributeValueType = iota
	// String of characters from the
	// set [0..9] and [A..F] that is prefixed with 0x or 0X.  The maximum
	// length of a hexadecimal-sequence depends on its AttributeNames.
	HexadecimalSequence
	// String of characters from the
	// set [0..9] and '.' that expresses a non-negative floating-point
	// number in decimal positional notation.
	DecimalFloatingPoint
	// String of characters
	// from the set [0..9], '-', and '.' that expresses a signed
	// floating-point number in decimal positional notation.
	SignedDecimalFloatingPoint
	// String of characters that does not have line feed (0xA), carriage return (0xD), or double
	// quote (0x22).
	QuotedString
	// Character string from a set that is
	// explicitly defined by the AttributeName.  An enumerated-string
	// will never contain double quotes ("), commas (,), or whitespace.
	EnumeratedString
	//Two decimal-integers separated by the "x"
	// character.  The first integer is a horizontal pixel dimension
	// (width); the second is a vertical pixel dimension (height).
	DecimalResolution
)
*/

type AttributeValuePair struct {
	// An AttributeName is an unquoted string containing characters from the
	// set [A..Z], [0..9] and '-'.  Therefore, AttributeNames contain only
	// uppercase letters, not lowercase.
	AttributeName string
	Value         string
}

type HLSTag struct {
	TagName string
	Value   string
}

var (
	LineIsNotATag error = errors.New("Line is not a valid HLS tag")
)

func ParseHLSTag(line string) (HLSTag, error) {
	hlsTag := HLSTag{}
	if !isTag(line) {
		return hlsTag, LineIsNotATag
	}

	line = strings.TrimPrefix(line, "#")
	rp := 0
	for rp < len(line) && line[rp] != ':' {
		rp++
	}

	hlsTag.TagName = line[:rp]
	if rp == len(line) {
		return hlsTag, nil
	}
	hlsTag.Value = line[rp+1:]

	return hlsTag, nil
}

func ParseCSV(csv string) []string {
	valueList := make([]string, 0)

	n := 0
	for n < len(csv) {
		token, readBytes := getCSV_Token(csv, n)
		n += readBytes
		valueList = append(valueList, token)
	}
	return valueList
}

// getCSV_Token scans csv from readPosition and returns CSV and bytes read.
func getCSV_Token(csv string, readPosition int) (string, int) {
	if readPosition >= len(csv) || csv[readPosition] == ',' {
		return "", 1
	}

	var quote bool
	enp := readPosition //End position
	for enp < len(csv) {
		if csv[enp] == '"' {
			quote = !quote
		} else if csv[enp] == ',' && !quote {
			break
		}
		enp++
	}

	n := enp - readPosition
	if enp == len(csv) {
		return csv[readPosition:], n
	} else if n == 0 {
		return "", n
	}

	return csv[readPosition:enp], n + 1
}
