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
	panic("Unknown LineType")
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
