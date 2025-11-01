package HLS_test

import (
	"bufio"
	"io"
	"testing"

	"net/http"

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

func ExamplePlaylist() {
	//This function serves a trailer to a movie in a media-playlist file.
	http.HandleFunc("/watch/025252/dracula-trailer", func(w http.ResponseWriter, r *http.Request) {
		playlist := HLS.NewPlaylist()

		//SetHeader append tag EXTM3U and EXT_X_VERSION. 7 is the version.
		err := playlist.SetHeader(7)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		{
			if err := playlist.AppendTag(HLS.HLSTag{
				TagName: HLS.EXTINF,
				Value:   "10.1",
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/pt1.ts")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		{
			if err := playlist.AppendTag(HLS.HLSTag{
				TagName: HLS.EXTINF,
				Value:   "9",
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/pt2.ts")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		{
			if err := playlist.AppendTag(HLS.HLSTag{
				TagName: HLS.EXTINF,
				Value:   "30.5",
			}); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			if err := playlist.AppendLine(HLS.NewPlaylistToken(HLS.RelativeURI, "/pt3.ts")); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		// Playlist must be closed before start reading.
		playlist.Close()

		rd := bufio.NewReader(&playlist)
		rd.WriteTo(w)
	})

	// Serves MPEG-4(.ts) for dracula-trailer.
	// ...
}
