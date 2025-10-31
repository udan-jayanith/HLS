package HLS_test

import (
	"github.com/udan-jayanith/HLS"
	"maps"
	"testing"
)

func TestParseAttributeList(t *testing.T) {
	{
		_, err := HLS.ParseAttributeList("")
		if err == nil {
			t.Log("Expected a error but got no.")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList("= ")
		if err != HLS.InvalidAttributeList {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList(" =")
		if err != HLS.InvalidAttributeList {
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
			t.Fatal("Attributes list has unexpected amount of value. Expected length to 1 but it has", len(attributes))
		}
	}

	{
		_, err := HLS.ParseAttributeList("attribute=")
		if err != HLS.InvalidAttributeList {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList("=value")
		if err != HLS.InvalidAttributeList {
			t.Log("Error ParseAttributeList")
			t.Fatal(err)
		}
	}

	{
		_, err := HLS.ParseAttributeList("attribute = value")
		if err != HLS.InvalidAttributeList {
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
			"CODECS":            `"avc1.64002a"`,
			"RESOLUTION":        "1920x1080",
			"URI":               `"v7/iframe_index.m3u8"`,
		}

		if !maps.Equal(attributes, resMap) {
			t.Log("Expected")
			t.Log(resMap)
			t.Log("but got")
			t.Log(attributes)
			t.FailNow()
		} else if len(attributes) != len(resMap) {
			t.Fatal("Attributes contains unexpected key value pairs")
		}
	}
}

func TestAttributeListString(t *testing.T) {
	//
	input := `AVERAGE-BANDWIDTH=183689,BANDWIDTH=187492,CODECS="avc1.64002a",RESOLUTION=1920x1080,URI="v7/iframe_index.m3u8"`
	attributes, err := HLS.ParseAttributeList(input)
	if err != nil {
		t.Fatal(err)
	}

	output, err := HLS.ParseAttributeList(attributes.String())
	if err != nil {
		t.Fatal(err)
	}

	if !maps.Equal(output, attributes) {
		t.Log("Expected")
		t.Log(attributes)
		t.Log("but got")
		t.Log(output)
		t.FailNow()
	}
}
