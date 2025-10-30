package HLS

import (
	"io"
)

var (
	//BlankLine is a blank line PlaylistToken.
	BlankLine PlaylistToken = NewPlaylistToken(Blank, "")
)

// Serialize returns the string-representation(HTTP Live streaming line) of the PlaylistToken.
func (pt *PlaylistToken) Serialize() string {
	return string(pt.SerializeAsBytes())
}

// Serialize returns the HTTP Live streaming line.
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

// NewPlaylistToken returns a new PlaylistToken.
func NewPlaylistToken(lineType LineType, value string) PlaylistToken {
	return PlaylistToken{
		Type:  lineType,
		Value: value,
	}
}

// Playlist represents a HLS content.
type Playlist struct {
	buf []byte
	err error
}

// NewPlaylist returns a new Playlist.
func NewPlaylist() Playlist {
	return Playlist{
		buf: make([]byte, 0),
	}
}

// AppendLine appends a PlaylistToken to underlying buffer for reading.
func (p *Playlist) AppendLine(playlistToken PlaylistToken) error {
	if p.err != nil {
		return p.err
	}
	p.buf = append(p.buf, playlistToken.SerializeAsBytes()...)
	return p.err
}

// Close closes the “pl *Playlist“ for AppendLine and reading.
// Read wait for closing until all data have been read (buffer must be empty to Read to return io.EOF).
// If AppendLine gets executed after Closing AppendLine returns a error.
func (pl *Playlist) Close() error {
	if pl.err != nil {
		return pl.err
	}
	pl.err = io.EOF
	return nil
}

// Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered.
// Even if Read returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.
// p[:n] data can be send to a clint for streaming.
func (pl *Playlist) Read(p []byte) (n int, err error) {
	if len(pl.buf) == 0 && pl.err == io.EOF {
		return n, pl.err
	} else if pl.err != nil && pl.err != io.EOF {
		return n, pl.err
	}
	return copy(p, pl.buf), nil
}
