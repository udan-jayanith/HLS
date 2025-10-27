package HLS

import (
	"errors"
	"strconv"
	"strings"
)

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

var (
	InvalidCSV error = errors.New("Invalid CSV")
)

func ParseCSV(csv string) ([]string, error) {
	tokens := make([]string, 0, 1)

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
			} else if token[0] == '"' && len(token)-1 != 0 && token[len(token)-1] == '"' {
				token = strings.Trim(token, `"`)
			}
			tokens = append(tokens, token)
		}
	}
	if quote {
		return tokens, InvalidCSV
	}

	return tokens, nil
}

var (
	InvalidAttributeList error = errors.New("Invalid attribute list")
)

// ParseAttributeList parses the attribute list and returns a map as attribute/value pair and a error.
// quoted-strings double quotes get removed.
func ParseAttributeList(attributeList string) (map[string]string, error) {
	csvs, err := ParseCSV(attributeList)
	attributeValuePairs := make(map[string]string, len(csvs))
	if err != nil {
		return attributeValuePairs, InvalidAttributeList
	}

	for _, csv := range csvs {
		rp := 0
		for rp < len(csv) && csv[rp] != '=' {
			rp++
		}

		//Returns a error if '=' sign is the last character of the csv or if csv[rp] is the last character or if csv[rp] sign is the first character.
		if rp >= len(csv)-1 || rp == 0 {
			return attributeValuePairs, InvalidAttributeList
		} else if csv[rp-1] == ' ' || csv[rp+1] == ' ' {
			return attributeValuePairs, InvalidAttributeList
		}
		attributeValuePairs[csv[:rp]] = strings.Trim(csv[rp+1:], `"`)
	}
	return attributeValuePairs, nil
}

type Resolution struct {
	Width, Height int
}

var (
	InvalidDecimalResolution error = errors.New("Invalid decimal resolution")
)

// decimalResolution: two decimal-integers separated by the "x"
// character.  The first integer is a horizontal pixel dimension
// (width); the second is a vertical pixel dimension (height).
//
// ParseResolution parses decimalResolution and returns Resolution
func ParseResolution(decimalResolution string) (Resolution, error) {
	//1024x720
	resolution := Resolution{}
	if !IsDecimalResolution(decimalResolution) {
		return resolution, InvalidDecimalResolution
	}

	var i int
	for i < len(decimalResolution) && decimalResolution[i] != 'x' {
		i++
	}
	if i+1 >= len(decimalResolution) || i == 0 {
		return resolution, InvalidDecimalResolution
	}

	width, err := strconv.Atoi(decimalResolution[:i])
	if err != nil {
		return resolution, InvalidDecimalResolution
	}

	height, err := strconv.Atoi(decimalResolution[i+1:])
	if err != nil {
		return resolution, InvalidDecimalResolution
	}

	resolution.Width = width
	resolution.Height = height
	return resolution, nil
}
