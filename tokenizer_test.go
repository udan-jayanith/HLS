package HLS_test

import (
	"HLS"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPlaylistTokenizer(t *testing.T) {
	file, err := os.Open("./example-m3u8/master.m3u8")
	if err != nil {
		t.Log("File opening error")
		t.Fatal(err)
	}

	playlistTokenizer := HLS.NewPlayListTokenizer(file)
	for {
		token, err := playlistTokenizer.Advance()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Log("Tokenizer Advanced() error")
			t.Fatal(err)
		} else if strings.TrimSpace(token.Value) == "" {
			t.Log("Blank line returned by Advanced()")
			t.Fatal(err)
		}
	}
}
