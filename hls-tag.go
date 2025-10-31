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

func (al *AttributeList) String() string {
	return ""
}

func (al AttributeList) Close() error {
	clear(al)
	return nil
}
