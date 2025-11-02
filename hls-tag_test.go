package HLS_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/udan-jayanith/HLS"
)

func TestParseHLSTag(t *testing.T) {
	//Test 1
	{
		hlsTag, err := HLS.ParseHLSTag("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=288000,RESOLUTION=256x144")
		if err != nil {
			t.Log("Failed parse ParseHLSTag")
			t.Fatal(err)
		} else if hlsTag.TagName != "EXT-X-STREAM-INF" {
			t.Fatal("Incorrect tag name. Expected", "EXT-X-STREAM-INF", "but got", hlsTag.TagName)
		} else if hlsTag.Value != "PROGRAM-ID=1,BANDWIDTH=288000,RESOLUTION=256x144" {
			t.Fatal("Unexpected value. Expected", "PROGRAM-ID=1,BANDWIDTH=288000,RESOLUTION=256x144", "but got", hlsTag.Value)
		}
	}

	//Test 2
	{
		hlsTag, err := HLS.ParseHLSTag("#EXTM3U")
		if err != nil {
			t.Log("Failed parse ParseHLSTag")
			t.Fatal(err)
		} else if hlsTag.TagName != "EXTM3U" {
			t.Fatal("Incorrect tag name. Expected", "EXTM3U", "but got", hlsTag.TagName)
		} else if hlsTag.Value != "" {
			t.Fatal("Unexpected value. Expected a empty string", "but got", hlsTag.Value)
		}
	}

	//Test 3
	{
		hlsTag, err := HLS.ParseHLSTag("#EXTM3U:")
		if err != nil {
			t.Log("Failed parse ParseHLSTag")
			t.Fatal(err)
		} else if hlsTag.TagName != "EXTM3U" {
			t.Fatal("Incorrect tag name. Expected", "EXTM3U", "but got", hlsTag.TagName)
		} else if hlsTag.Value != "" {
			t.Fatal("Unexpected value. Expected a empty string", "but got", hlsTag.Value)
		}
	}

	//Test 4
	{
		if _, err := HLS.ParseHLSTag("4k/skate_phantom_flex_4k_8288_2160p.m3u8"); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}

	//Test 5
	{
		if _, err := HLS.ParseHLSTag("#This is a comment"); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}

	//Test 5
	{
		if _, err := HLS.ParseHLSTag(""); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}
}

func TestToPlaylistToken(t *testing.T) {
	hlsTag := HLS.HLSTag{
		TagName: HLS.EXTINF,
		Value:   `21.3,"title"`,
	}

	playlistToken := hlsTag.ToPlaylistToken()
	if playlistToken.Type != HLS.Tag {
		t.Fatal("Expected", HLS.Tag.String(), "but got", playlistToken.Type.String())
	} else if playlistToken.Serialize() != `#EXTINF:21.3,"title"`+"\n" {
		t.Fatal("Expected", `#EXTINF:21.3,"title"`, "but got", playlistToken.Serialize())
	}
}

func ExampleParseHLSTag() {
	hlsTag, err := HLS.ParseHLSTag("#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=288000,RESOLUTION=256x144")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hlsTag.TagName)
	fmt.Println(hlsTag.Value)
	//Output:
	//EXT-X-STREAM-INF
	//PROGRAM-ID=1,BANDWIDTH=288000,RESOLUTION=256x144
}
