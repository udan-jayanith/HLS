package HLS_test

import (
	"HLS"
	"slices"
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	{
		list := []string{
			"PROGRAM-ID=1",
			"BANDWIDTH=10768000",
			`CODECS="avc1.640028,mp4a.40.2"`,
			"RESOLUTION=3840x2160",
		}
		csv := strings.Join(list, ",")
		values, err := HLS.ParseCSV(csv)
		if err != nil {
			t.Log(values)
			t.Fatal(err)
		}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		list := []string{
			"PROGRAM-ID=1",
			"BANDWIDTH=10768000",
			`CODECS="avc1.640028,mp4a.40.2",`,
			"RESOLUTION=3840x2160",
		}
		csv := strings.Join(list, ",")
		if _, err := HLS.ParseCSV(csv); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}

	{
		values, err := HLS.ParseCSV(`
			Enumerated-String,1240x720,"Quoted String"
		`)
		if err != nil {
			t.Fatal(err)
		}
		list := []string{
			"Enumerated-String",
			"1240x720",
			`"Quoted String"`,
		}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		if _, err := HLS.ParseCSV(", "); err == nil {
			t.Fatal("Expected a error but got no error")
		}

	}

	{
		if _, err := HLS.ParseCSV(" ,"); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}

	{
		values, err := HLS.ParseCSV("value,")
		if err != nil {
			t.Log("Unexpected error")
			t.Fatal(err)
		}
		list := []string{
			"value",
		}
		if slices.Compare(values, list) != 0 {
			t.Fatal("Expected", list, "but got", values)
		}
	}

	{
		if _, err := HLS.ParseCSV(","); err == nil {
			t.Fatal("Expected a error but got no error")
		}
	}

	{
		if _, err := HLS.ParseCSV(""); err == nil {
			t.Fatal("Expected a error but got no error")
		}

	}
}

func TestCSV_String(t *testing.T) {
	input := "<time>,[<title>]"
	csvs, err := HLS.ParseCSV(input)
	if err != nil {
		t.Fatal(err)
	}
	if csvs.String() != input {
		t.Fatal("Expected", input, "but got", csvs.String())
	}
}
