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
	Blank
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
	case Blank:
		return "Blank"
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
	return Blank
}

type PlayListTokenizer struct {
	rd       *bufio.Reader
	eofError error
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

// Advanced read from plt.rd(hls playlist) until '\n' and returns a Token and a error. Advanced return an error io.EOF if reading is finished.
// If error is not nil or io.EOF hls file is broken.
// Advanced does not returns blank line types.
// If token type is a Comment or a Tag '#' prefix get removed.
func (plt *PlayListTokenizer) Advance() (PlaylistToken, error) {
	for {
		token := PlaylistToken{}
		if plt.eofError != nil {
			return token, plt.eofError
		}

		line, err := plt.rd.ReadString('\n')
		if err == io.EOF {
			plt.eofError = err
		} else if err != nil {
			return token, err
		}

		line = strings.TrimSpace(line)
		token.Type = getLineType(line)

		if token.Type == Comment || token.Type == Tag {
			token.Value = strings.TrimPrefix(line, "#")
		} else {
			token.Value = line
		}
		if token.Type == Tag {
			token.Value = strings.TrimSuffix(token.Value, ",")
		}

		if token.Type != Blank {
			return token, nil
		}
	}
}
