package HLS_test

import (
	"HLS"
	"testing"
)

func TestIsDecimalInteger(t *testing.T) {
	{
		if HLS.IsDecimalInteger("184467440737095516150") {
			t.Fatal("Expected to return false")
		}
	}
	{
		if HLS.IsDecimalInteger(`"16363"`) {
			t.Fatal("Expected to return false")
		}
	}
	{
		if HLS.IsDecimalInteger(``) {
			t.Fatal("Expected to return false")
		}
	}
	{
		if HLS.IsDecimalInteger(` `) {
			t.Fatal("Expected to return false")
		}
	}
	{
		if !HLS.IsDecimalInteger(`1`) {
			t.Fatal("Expected to return true")
		}
	}
	{
		if HLS.IsDecimalInteger(`-1`) {
			t.Fatal("Expected to return false")
		}
	}
	{
		if !HLS.IsDecimalInteger(`10`) {
			t.Fatal("Expected to return true")
		}
	}
	{
		if !HLS.IsDecimalInteger(`18446744073709551615`) {
			t.Fatal("Expected to return true")
		}
	}
}

func TestIsHexadecimalSequence(t *testing.T) {
	{
		if HLS.IsHexadecimalSequence("") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsHexadecimalSequence(" ") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsHexadecimalSequence("ABC-123") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsHexadecimalSequence("abc123") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if !HLS.IsHexadecimalSequence("ABC123") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsHexadecimalSequence("123ABC") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsHexadecimalSequence("1ABC23") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsHexadecimalSequence("ABC") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsHexadecimalSequence("123") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if HLS.IsHexadecimalSequence("1.23") {
			t.Fatal("Expected to be false")
		}
	}
}

func TestDecimalFloatingPoint(t *testing.T) {
	{
		if HLS.IsDecimalFloatingPoint("") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsDecimalFloatingPoint(" ") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsDecimalFloatingPoint("14.42.52") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsDecimalFloatingPoint("ABC") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsDecimalFloatingPoint(".10") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsDecimalFloatingPoint("10.") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if !HLS.IsDecimalFloatingPoint("10") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsDecimalFloatingPoint("0.1031") {
			t.Fatal("Expected to be true")
		}
	}
}

func TestSignedDecimalFloatingPoint(t *testing.T) {
	{
		if HLS.IsSignedDecimalFloatingPoint("") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsSignedDecimalFloatingPoint(" ") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsSignedDecimalFloatingPoint("12-41") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsSignedDecimalFloatingPoint("42-.4") {
			t.Fatal("Expected to be false")
		}
	}
	{
		if HLS.IsSignedDecimalFloatingPoint("4-") {
			t.Fatal("Expected to be false")
		}
	}

	{
		if !HLS.IsSignedDecimalFloatingPoint("-4.5") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsSignedDecimalFloatingPoint("4") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsSignedDecimalFloatingPoint("-4") {
			t.Fatal("Expected to be true")
		}
	}
	{
		if !HLS.IsSignedDecimalFloatingPoint("31.4") {
			t.Fatal("Expected to be true")
		}
	}
}

func TestIsString(t *testing.T) {
	testcases := []struct {
		value  string
		output bool
	}{
		{
			value: "\r\n",
		},
		{
			value: "\n",
		},
		{
			value: `""`,
		},
		{
			value:  "",
			output: true,
		}, {
			value:  " ",
			output: true,
		},
		{
			value:  "Hello world",
			output: true,
		},
	}
	for _, testcase := range testcases {
		output := HLS.IsString(testcase.value)
		if output != testcase.output {
			t.Log(testcase)
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}
}

func TestQuotedString(t *testing.T) {
	type testcase struct {
		value  string
		output bool
	}

	testcases := []testcase{
		{
			value:  "",
			output: false,
		},
		{
			value:  " ",
			output: false,
		},
		{
			value:  `"Hello world`,
			output: false,
		},
		{
			value:  `Hello world"`,
			output: false,
		},
		{
			value: `"`,
		},
		{
			value:  `""`,
			output: true,
		},
		{
			value:  `"Hello world"`,
			output: true,
		},
	}

	for _, testcase := range testcases {
		output := HLS.IsQuotedString(testcase.value)
		if output != testcase.output {
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}
}

func TestIsEnumeratedString(t *testing.T) {
	//""
	//" "
	//Hello,
	//Hello" world
	//Hello world
	type testcase struct {
		value  string
		output bool
	}

	testcases := []testcase{
		{
			value: "",
		},
		{
			value: " ",
		},
		{
			value: `Hello,`,
		},
		{
			value: `"Hello-world"`,
		},
		{
			value: `Hello world`,
		},
		{
			value:  `Hello-world`,
			output: true,
		},
	}

	for _, testcase := range testcases {
		output := HLS.IsEnumeratedString(testcase.value)
		if output != testcase.output {
			t.Log("input", testcase.value)
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}

}

func TestIsDecimalResolution(t *testing.T) {
	//""
	//" "
	//1024x
	//x720
	//1024
	//1024x260x720
	//1024x720abc

	//1024x720
	//0x0
	type testcase struct {
		value  string
		output bool
	}

	testcases := []testcase{
		{
			value: "",
		},
		{
			value: " ",
		},
		{
			value: `1024x`,
		},
		{
			value: `x720`,
		},
		{
			value: `1024`,
		},
		{
			value: `1024x260x720`,
		},
		{
			value: "1024x720abc",
		},
		{
			value:  "1024x720",
			output: true,
		},
		{
			value:  "0x0",
			output: true,
		},
	}

	for _, testcase := range testcases {
		output := HLS.IsDecimalResolution(testcase.value)
		if output != testcase.output {
			t.Log("input", testcase.value)
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}
}

func TestWrapQuotes(t *testing.T) {
	//""
	//" "
	//"Hello world"
	testcases := []struct {
		value  string
		output string
	}{
		{
			value:  "",
			output: `""`,
		},
		{
			value:  " ",
			output: `" "`,
		},
		{
			value:  `Hello world`,
			output: `"Hello world"`,
		},
	}
	for _, testcase := range testcases {
		output := HLS.WrapQuotes(testcase.value)
		if output != testcase.output {
			t.Log(testcase)
			t.Fatal("Expected", testcase.output, "but got", output)
		}
	}

}

func TestParseDecimalResolution(t *testing.T) {
	{
		resolution, err := HLS.ParseResolution("1024x720")
		if err != nil {
			t.Log("Unexpected error")
			t.Fatal(err)
		} else if resolution.Width != 1024 {
			t.Fatal("Expected width of 1024 but got", resolution.Width)
		} else if resolution.Height != 720 {
			t.Fatal("Expected height of 720 but got", resolution.Height)
		}
	}

	{
		if _, err := HLS.ParseResolution(""); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}

	{
		if _, err := HLS.ParseResolution(" "); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}

	{
		if _, err := HLS.ParseResolution("1024x"); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}

	{
		if _, err := HLS.ParseResolution("x720"); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}

	{
		if _, err := HLS.ParseResolution("x"); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}

	{
		if _, err := HLS.ParseResolution("a"); err != HLS.InvalidDecimalResolution {
			t.Fatal("Expected error", HLS.InvalidDecimalResolution, "but got", err)
		}
	}
}

func TestToDecimalResolution(t *testing.T) {
	{
		resolution := HLS.Resolution{}
		output := resolution.ToDecimalResolution()
		if output != "0x0" {
			t.Fatal("Expected 0x0 but got", output)
		}
	}

	{
		resolution := HLS.Resolution{
			Width:  1024,
			Height: 720,
		}
		output := resolution.ToDecimalResolution()
		if output != "1024x720" {
			t.Fatal("Expected 1024x720 but got", output)
		}
	}
}
