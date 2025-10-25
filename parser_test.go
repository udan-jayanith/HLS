package HLS_test

import (
	"HLS"
	"slices"
	"strings"
	"testing"
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

func TestParseCSV(t *testing.T) {
	{
		list := []string{
			"PROGRAM-ID=1",
			"BANDWIDTH=10768000",
			`CODECS="avc1.640028,mp4a.40.2"`,
			"RESOLUTION=3840x2160",
		}
		csv := strings.Join(list, ",")
		values := HLS.ParseCSV(csv)
		if slices.Compare(list, values) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		list := []string{
			"PROGRAM-ID=1",
			"BANDWIDTH=10768000",
			`CODECS="avc1.640028,mp4a.40.2"`,
			"RESOLUTION=3840x2160",
		}
		csv := strings.Join(list, ",") + ","
		values := HLS.ParseCSV(csv)
		if slices.Compare(list, values) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		values := HLS.ParseCSV(", ")
		if slices.Compare(values, []string{
			"",
			" ",
		}) != 0 {
			t.Fatal("Expected two empty strings but got", values)
		}
	}

	{
		values := HLS.ParseCSV(" ,")
		if slices.Compare(values, []string{
			" ",
		}) != 0 {
			t.Fatal("Expected ` ,` but got", values)
		}
	}

	{
		values := HLS.ParseCSV("value,")
		list := []string{
			"value",
		}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		values := HLS.ParseCSV(",")
		list := []string{
			"",
		}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		values := HLS.ParseCSV("")
		list := []string{}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}
}
