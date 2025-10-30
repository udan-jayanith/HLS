package HLS

import (
	"errors"
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

// Returned []string value is in left to right as in csv string.
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
			}
			tokens = append(tokens, token)
		}
	}
	if quote {
		return tokens, InvalidCSV
	}

	return tokens, nil
}

type CSV_Builder struct {
	builder strings.Builder
}

func (cb *CSV_Builder) Append(value ...string) {

}

func (cb *CSV_Builder) String() string {
	return ""
}

var (
	InvalidAttributeList error = errors.New("Invalid attribute list")
)

type AttributeList map[string]string

// ParseAttributeList parses the attribute list and returns a map as attribute/value pair and a error.
func ParseAttributeList(attributeList string) (AttributeList, error) {
	csvs, err := ParseCSV(attributeList)
	attributeValuePairs := make(AttributeList, len(csvs))
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
		attributeValuePairs[csv[:rp]] = csv[rp+1:]
	}
	return attributeValuePairs, nil
}

func (al *AttributeList) String() string{
	return ""
}
