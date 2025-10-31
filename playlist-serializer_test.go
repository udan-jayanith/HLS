package HLS_test

import (
	"github.com/udan-jayanith/HLS"
	"bufio"
	"io"
	"testing"
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