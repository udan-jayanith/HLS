package HLS

import "strings"

//decimal-integer: an unquoted string of characters from the set
//[0..9] expressing an integer in base-10 arithmetic in the range
//from 0 to 2^64-1 (18446744073709551615).  A decimal-integer may be
//from 1 to 20 characters long.
func IsDecimalInteger(value string) bool {
	if len(value) < 1 || len(value) > 20 {
		return false
	}
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

//hexadecimal-sequence: an unquoted string of characters from the
//set [0..9] and [A..F].  The maximum
//length of a hexadecimal-sequence depends on its AttributeNames.
func IsHexadecimalSequence(value string) bool {
	for _, char := range value {
		if !(char >= '0' && char <= '9') && !(char >= 'A' && char <= 'Z') {
			return false
		}
	}
	return true
}

//decimal-floating-point: an unquoted string of characters from the
//set [0..9] and '.' that expresses a non-negative floating-point
//number in decimal positional notation.
func IsDecimalFloatingPoint(value string) bool {
	for _, char := range value {
		if (char < '0' || char > '9') && char != '.' {
			return false
		}
	}
	return true
}

//signed-decimal-floating-point: an unquoted string of characters
//from the set [0..9], '-', and '.' that expresses a signed
//floating-point number in decimal positional notation.
func IsSignedDecimalFloatingPoint(value string) bool {
	for _, char := range value {
		if (char < '0' || char > '9') && char != '.' && char != '-' {
			return false
		}
	}
	return true
}

//The following characters MUST NOT appear in a
//string: line feed (0xA), carriage return (0xD), or double
//quote (0x22).
func IsString(value string) bool {
	//\n, \r, "
	for _, char := range value {
		if char == '\n' || char == '\r' || char == '"' {
			return false
		}
	}
	return true
}

//quoted-string: a string of characters within a pair of double
//quotes (0x22).  The following characters MUST NOT appear in a
//quoted-string: line feed (0xA), carriage return (0xD), or double
//quote (0x22).  Quoted-string AttributeValues SHOULD be constructed
//so that byte-wise comparison is sufficient to test two quoted-
//string AttributeValues for equality.  Note that this implies case-
//sensitive comparison.
func IsQuotedString(value string) bool {
	if len(value) < 2 {
		return false
	} else if value[0] != '"' || value[len(value)-1] != '"' {
		return false
	}
	return IsString(strings.Trim(value, `"`))
}

//enumerated-string: an unquoted character string from a set that is
//explicitly defined by the AttributeName.  An enumerated-string
//will never contain double quotes ("), commas (,), or whitespace.
func EnumeratedString(value string) bool {
	for _, char := range value {
		if char == ',' || char == ' ' || char == '"' {
			return false
		}
	}
	return true
}

//decimal-resolution: two decimal-integers separated by the "x"
//character.  The first integer is a horizontal pixel dimension
//(width); the second is a vertical pixel dimension (height).
func IsDecimalResolution(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}

	values := strings.SplitN(value, "x", 1)
	if len(values) != 2 {
		return false
	}

	return IsDecimalInteger(values[0]) && IsDecimalInteger(values[1])
}

func WrapQuotes(value string) string {
	return `"` + value + `"`
}
