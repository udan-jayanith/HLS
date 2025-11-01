package HLS

import (
	"errors"
	"fmt"
	"strings"
)

// HLSTag tag is used to represent a hls tag line where TagName is the tag name without '#' sign.
//
// Example:
//
//	EXTM3U
//
// Value is the comma separated values of the tag line if any.
type HLSTag struct {
	TagName PlaylistTag
	Value   string
}

var (
	LineIsNotATag error = errors.New("Line is not a valid HLS tag")
)

// ParseHLSTag parses a HLS tag line.
// Examples of line
//
//	#EXTINF:9.009,
//
// "#" at the beginning of the tag is optional.
func ParseHLSTag(line string) (HLSTag, error) {
	hlsTag := HLSTag{}
	line = strings.TrimPrefix(line, "#")
	if !isTag("#" + line) {
		return hlsTag, LineIsNotATag
	}

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

// ToPlaylistToken returns the PlaylistToken token of the HLSTag
func (ht *HLSTag) ToPlaylistToken() PlaylistToken {
	return PlaylistToken{
		Type: Tag,
		Value: func(ht *HLSTag) string {
			if strings.TrimSpace(ht.Value) == "" {
				return ht.TagName
			}
			return fmt.Sprintf("%s:%s", ht.TagName, ht.Value)
		}(ht),
	}
}
