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

	//Hello-world
}

func TestIsDecimalResolution(t *testing.T) {
	//""
	//" "
	//1024x
	//x720
	//1024x720
	//1024
	//0x0
	//1024x260x720
	//1024x720abc
}

func TestWrapQuotes(t *testing.T) {
	//""
	//" "
	//"Hello world"
}
