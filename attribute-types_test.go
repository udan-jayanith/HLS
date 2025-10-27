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
	//""
	//" "
	//ABC-123
	//abc123
	//ABC123
	//123ABC
	//1ABC23
}

func TestDecimalFloatingPoint(t *testing.T) {
	//""
	//" "
	//10
	//0.1031
	//14.42.52
}

func TestSignedDecimalFloatingPoint(t *testing.T) {
	//""
	//" "
	//12-41
	//31.4
	//42-.4
	//-4.5
	//4
	//-4
	//4-
}

func TestIsString(t *testing.T) {
	//\r\n
	//\n
	//"
	//""
	//" "
	//`""`
	//Hello world
}

func TestQuotedString(t *testing.T) {
	//""
	//" "
	//`""`
	//`"Hello world"`
	//`"Hello world`
	//`hello world"`
}

func TestIsEnumeratedString(t *testing.T) {
	//""
	//" "
	//Hello,
	//Hello world
	//Hello" world
	//Hello world
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
