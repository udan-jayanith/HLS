package HLS

import (
	"io"
)

const (
	BlankLine string = "\n"
)

func (pt *PlaylistToken) Serialize() string {
	return string(pt.SerializeAsBytes())
}

func (pt *PlaylistToken) SerializeAsBytes() []byte {
	res := []byte(pt.Value)
	switch pt.Type {
	case Comment, Tag:
		buf := make([]byte, 0, len(res)+1)
		buf = append(buf, '#')
		res = append(buf, res...)
	}
	return append(res, '\n')
}

func NewPlaylistToken(lineType LineType, value string) PlaylistToken {
	return PlaylistToken{
		Type:  lineType,
		Value: value,
	}
}

type Playlist struct {
	buf []byte
	err error
}

func NewPlaylist() Playlist {
	return Playlist{
		buf: make([]byte, 0),
	}
}

func (p *Playlist) AppendLine(playlistToken PlaylistToken) error {
	if p.err != nil {
		return p.err
	}
	p.buf = append(p.buf, playlistToken.SerializeAsBytes()...)
	return p.err
}

func (pl *Playlist) Close() error {
	if pl.err != nil {
		return pl.err
	}
	pl.err = io.EOF
	return nil
}

func (pl *Playlist) Read(p []byte) (n int, err error) {
	if len(pl.buf) == 0 && pl.err == io.EOF {
		return n, pl.err
	} else if pl.err != nil && pl.err != io.EOF {
		return n, pl.err
	}
	return copy(p, pl.buf), nil
}
