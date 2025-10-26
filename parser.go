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

var (
	InvalidAttributeValuePair error = errors.New("Invalid attribute value pair")
	ContainsInvalidSpaces     error = errors.New("Contains invalid spaces")
)

// ParseAttributeList parses the attribute list and returns a map as attribute/value pair and a error.
// quoted-strings double quotes get removed.
func ParseAttributeList(attributeList string) (map[string]string, error) {
	csvs := ParseCSV(attributeList)
	attributeValuePairs := make(map[string]string, len(csvs))

	for _, csv := range csvs {
		rp := 0
		for rp < len(csv) && csv[rp] != '=' {
			rp++
		}

		//Returns a error if '=' sign is the last character of the csv or if csv[rp] is the last character or if csv[rp] sign is the first character.
		if rp >= len(csv)-1 || rp == 0 {
			return attributeValuePairs, InvalidAttributeValuePair
		} else if csv[rp-1] == ' ' || csv[rp+1] == ' ' {
			return attributeValuePairs, ContainsInvalidSpaces
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
	if len(decimalResolution) == 0 || decimalResolution[0] == 'x' || decimalResolution[len(decimalResolution)-1] == 'x' {
		return Resolution{}, InvalidDecimalResolution
	}

	xPosition := 0
	for xPosition < len(decimalResolution) && decimalResolution[xPosition] != 'x' {
		xPosition++
	}
	if xPosition >= len(decimalResolution)-1 {
		return Resolution{}, InvalidDecimalResolution
	}

	width, err := strconv.Atoi(decimalResolution[:xPosition])
	if err != nil {
		return Resolution{}, err
	}

	height, err := strconv.Atoi(decimalResolution[xPosition+1:])
	if err != nil {
		return Resolution{}, err
	}

	return Resolution{
		Width:  width,
		Height: height,
	}, nil
}
