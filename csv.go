package HLS

import (
	"errors"
	"strings"
)

var (
	InvalidCSV error = errors.New("Invalid CSV")
)

type csvs []string

// Returned []string value is in left to right as in csv string.
func ParseCSV(csv string) (csvs, error) {
	tokens := make(csvs, 0, 1)

	if len(csv) < 1 {
		return tokens, InvalidCSV
	} else if csv[len(csv)-1] != ',' {
		csv += string(',')
	}
	var quote bool
	var comma int
	for i, char := range csv {
		if char == '"' {
			quote = !quote
		} else if !quote && char == ' ' {
			return tokens, InvalidCSV
		} else if !quote && char == ',' {
			token := strings.TrimSpace(csv[comma:i])
			comma = i + 1
			if len(token) < 1 {
				return tokens, InvalidCSV
			}
			tokens = append(tokens, token)
		}
	}
	if quote {
		return tokens, InvalidCSV
	}

	return tokens, nil
}

func (cb csvs) String() string {
	return strings.Join(cb, ",")
}
