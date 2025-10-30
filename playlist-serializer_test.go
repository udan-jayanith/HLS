package HLS_test

import (
	"HLS"
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

func TestPlaylist(t *testing.T) {}
