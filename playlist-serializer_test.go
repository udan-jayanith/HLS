package HLS_test

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/udan-jayanith/HLS"
)

func TestPlaylistTokenSerialize(t *testing.T) {
	testcases := []struct {
		token  HLS.PlaylistToken
		output string
	}{
		{
			token:  HLS.NewPlaylistToken(HLS.Tag, HLS.EXTM3U),
			output: "#EXTM3U\n",
		},
		{
			token:  HLS.NewPlaylistToken(HLS.Comment, "This is a comment"),
			output: "#This is a comment\n",
		},
		{
			token:  HLS.NewPlaylistToken(HLS.URI, "http://media.example.com/first.ts"),
			output: "http://media.example.com/first.ts\n",
		},
		{
			token:  HLS.NewPlaylistToken(HLS.RelativeURI, "/third.ts"),
			output: "/third.ts\n",
		},
	}

	for _, testcase := range testcases {
		output := testcase.token.Serialize()
		if output != testcase.output {
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}

	{
		output := HLS.BlankLine.Serialize()
		if HLS.BlankLine.Serialize() != "\n" {
			t.Fatal("Expected newline character but got", output)
		}
	}
}

func TestPlaylist(t *testing.T) {
	playlist := HLS.NewPlaylist()

	playlist.AppendLine(HLS.NewPlaylistToken(HLS.Tag, HLS.EXTM3U))
	playlist.AppendLine(HLS.NewPlaylistToken(HLS.Comment, "This is a comment"))
	playlist.AppendLine(HLS.BlankLine)
	playlist.AppendLine(HLS.NewPlaylistToken(HLS.URI, "http://media.example.com/first.ts"))
	playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/third.ts"))

	if err := playlist.Close(); err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		output string
		err    error
	}{
		{
			output: "#EXTM3U\n",
		},
		{
			output: "#This is a comment\n",
		},
		{
			output: "\n",
		},
		{
			output: "http://media.example.com/first.ts\n",
		},
		{
			output: "/third.ts\n",
		},
		{
			output: "",
			err:    io.EOF,
		},
	}

	rd := bufio.NewReader(&playlist)
	for _, testcase := range testcases {
		line, err := rd.ReadString('\n')
		if err != testcase.err {
			t.Fatal("Expected", testcase.err, "but got", err)
		} else if line != testcase.output {
			t.Fatal("Expected", testcase.output, "but got", line)
		}
	}

}

func TestPlaylist_SetHeader(t *testing.T) {
	playlist := HLS.NewPlaylist()

	err := playlist.SetHeader(7)
	if err != nil {
		t.Fatal(err)
	}

	tokenizer := HLS.NewPlayListTokenizer(&playlist)

	if token, err := tokenizer.Advance(); err != nil {
		t.Fatal(err)
	} else if token.Type != HLS.Tag {
		t.Fatal("Expected", HLS.Tag.String(), "but got", token.Type.String())
	} else if token.Value != "EXTM3U" {
		t.Fatal("Expected EXTM3U but got", token.Value)
	}

	if token, err := tokenizer.Advance(); err != nil {
		t.Fatal(err)
	} else if token.Type != HLS.Tag {
		t.Fatal("Expected", HLS.Tag.String(), "but got", token.Type.String())
	} else if token.Value != "EXT-X-VERSION:7" {
		t.Fatal("Expected EXT_X_VERSION:7 but got", token.Value)
	}

	if _, err := tokenizer.Advance(); err == nil {
		t.Fatal("Expected error io.EOF but got nil")
	}
}

// Returns the media-playlist.
func fullVideoMediaPlaylist() (string, error) {
	playlist := HLS.NewPlaylist()

	//SetHeader append tag EXTM3U and EXT_X_VERSION. 7 is the version.
	if err := playlist.SetHeader(7); err != nil {
		return "", err
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXT_X_TARGETDURATION,
			Value:   "14",
		}); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "11.266667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg000.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "13.766667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg001.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "7.166667",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg002.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "8.533333",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg003.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "11.800000",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg004.ts")); err != nil {
			return "", err
		}
	}

	{
		if err := playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXTINF,
			Value:   "7.433333",
		}); err != nil {
			return "", err
		}

		if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/seg005.ts")); err != nil {
			return "", err
		}
	}

	{
		playlist.AppendTag(HLS.HLSTag{
			TagName: HLS.EXT_X_ENDLIST,
		})
	}
	//	Close must be called before reading the playlist
	playlist.Close()

	var builder strings.Builder
	rd := bufio.NewReader(&playlist)
	rd.WriteTo(&builder)

	return builder.String(), nil
}

func ExamplePlaylist() {
	mediaPlaylist, err := fullVideoMediaPlaylist()
	if err != nil {
		log.Fatal(err)
	}

	// This function serves video segments.
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler := http.StripPrefix("/", http.FileServer(http.Dir("./video-fragments")))
		handler.ServeHTTP(w, r)
	}))

	// This function serves a trailer to a movie in a media-playlist file.
	http.HandleFunc("/full-video.m3u8", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mediaPlaylist))
	})

	//...
}
