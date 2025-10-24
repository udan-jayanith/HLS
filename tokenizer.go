package HLS

import (
	"bufio"
	"io"
	"strings"
)

type LineType int

const (
	Tag LineType = iota //#EXT
	URI
	RelativeURI
	Comment //#
	blank
)

func (t LineType) String() string {
	switch t {
	case Tag:
		return "Tag"
	case URI:
		return "URI"
	case RelativeURI:
		return "Relative URI"
	case Comment:
		return "Comment"
	case blank:
		return "blank"
	}
	return "Unknown LineType"
}

func isTag(line string) bool {
	return strings.HasPrefix(line, "#EXT")
}

func isURI(line string) bool {
	//http://, https://
	return (strings.HasPrefix(line, "https://") || strings.HasPrefix(line, "http://")) && !strings.HasPrefix(line, "#")
}

func isRelativeURI(line string) bool {
	return !isURI(line) && !strings.HasPrefix(line, "#") && line != ""
}

func isComment(line string) bool {
	return strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "#EXT")
}

func getLineType(line string) LineType {
	if isTag(line) {
		return Tag
	} else if isURI(line) {
		return URI
	} else if isRelativeURI(line) {
		return RelativeURI
	} else if isComment(line) {
		return Comment
	}
	return blank
}

type PlayListTokenizer struct {
	rd *bufio.Reader
}

// NewPlayListTokenizer returns a new PlayListTokenizer
func NewPlayListTokenizer(r io.Reader) PlayListTokenizer {
	return PlayListTokenizer{
		rd: bufio.NewReader(r),
	}
}

type PlaylistToken struct {
	Type  LineType
	Value string
}

// Advanced read from plt.rd until '\n' and returns a Token and a error. Advanced return an error io.EOF if reading is finished.
// If error is not nil or io.EOF hls file is broken.
// Advanced does not returns blank lines.
func (plt *PlayListTokenizer) Advanced() (PlaylistToken, error) {
	for {
		token := PlaylistToken{}
		line, err := plt.rd.ReadString('\n')
		if err != nil {
			return token, err
		}

		line = strings.TrimSpace(line)
		token.Type = getLineType(line)
		token.Value = line

		if token.Type != blank {
			return token, nil
		}
	}
}

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

type AttributeValuePair struct {
	Type          AttributeValueType
	AttributeName string
	Value         string
}

type HLSTag struct {
	TagName string
	// Certain tags have values that are attribute-lists.  An attribute-list
	// is a comma-separated list of attribute/value pairs with no
	// whitespace.
	AttributeList []AttributeValuePair
}

/*
type AttributeListTokenizer struct {
	value string
}

func NewAttributeListTokenizer(tagValue string) AttributeListTokenizer {
	return AttributeListTokenizer{
		value: tagValue,
	}
}

func (alt *AttributeListTokenizer) Advanced() (string, error) {

}

*/
