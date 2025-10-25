package HLS_test

import (
	"HLS"
	"log"
	"maps"
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

func TestParseAttributeList(t *testing.T) {
	{
		attributes, err := HLS.ParseAttributeList("")
		if err != nil {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		} else if len(attributes) != 0 {
			t.Fatal("Expected attributes to be empty but not.")
		}
	}

	{
		_, err := HLS.ParseAttributeList("= ")
		if err != HLS.InvalidAttributeValuePair {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList(" =")
		if err != HLS.InvalidAttributeValuePair {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		attributes, err := HLS.ParseAttributeList("1=1")
		if err != nil {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}

		if val, ok := attributes["1"]; !ok {
			t.Fatal("Value is not found for the key")
		} else if val != "1" {
			t.Fatal("Unexpected value for val. Expected `1` but got", val)
		}

		if len(attributes) != 1 {
			log.Fatal("Attributes list has unexpected amount of value. Expected length to 1 but it has", len(attributes))
		}
	}

	{
		_, err := HLS.ParseAttributeList("attribute=")
		if err != HLS.InvalidAttributeValuePair {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList("=value")
		if err != HLS.InvalidAttributeValuePair {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList("attribute = value")
		if err != HLS.ContainsInvalidSpaces {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		attributes, err := HLS.ParseAttributeList("attribute=value")
		if err != nil {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		} else if len(attributes) != 1 {
			t.Fatal("Attributes has unexpected length. Expected length to be 1 but it has", len(attributes))
		} else if attributes["attribute"] != "value" {
			t.Fatal("Expected attribute=='value' but attribute==", attributes["attribute"])
		}
	}

	{
		attributes, err := HLS.ParseAttributeList(`AVERAGE-BANDWIDTH=183689,BANDWIDTH=187492,CODECS="avc1.64002a",RESOLUTION=1920x1080,URI="v7/iframe_index.m3u8"`)
		if err != nil {
			t.Fatal(err)
		}

		resMap := map[string]string{
			"AVERAGE-BANDWIDTH": "183689",
			"BANDWIDTH":         "187492",
			"CODECS":            `avc1.64002a`,
			"RESOLUTION":        "1920x1080",
			"URI":               "v7/iframe_index.m3u8",
		}

		if !maps.Equal(attributes, resMap) {
			t.Log("Unexpected attributes")
			t.Log(resMap)
			t.Log("but got")
			t.Log(attributes)
			t.FailNow()
		} else if len(attributes) != len(resMap) {
			t.Fatal("Attributes contains unexpected key value pairs")
		}
	}

}
