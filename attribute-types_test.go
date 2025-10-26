package HLS_test

import "testing"

func TestIsDecimalInteger(t *testing.T) {
	//184467440737095516150
	//"16363"
	//""
	//" "
	//1
	//-1
	//10
	//18446744073709551615
}

func TestIsHexadecimalSequence(t *testing.T) {
	//""
	//" "
	//ABC123
	//ABC-123
	//abc123
	//123ABC
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
