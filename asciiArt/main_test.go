package asciiArt

import (
	_const "ASCIIArt/asciiArt/constants"
	"testing"
)

func TestConvertToAsciiArtWithStandard(t *testing.T) {
	got := СonvertToAsciiArt("A", _const.ResourcesPath+"standard.txt")
	want :=
		"           \n" +
			"    /\\     \n" +
			"   /  \\    \n" +
			"  / /\\ \\   \n" +
			" / ____ \\  \n" +
			"/_/    \\_\\ \n" +
			"           \n" +
			"           "
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}

func TestConvertToAsciiArtWithThinkertoy(t *testing.T) {
	got := СonvertToAsciiArt("A", _const.ResourcesPath+"thinkertoy.txt")
	want :=
		"      \n" +
			"  O   \n" +
			" / \\  \n" +
			"o---o \n" +
			"|   | \n" +
			"o   o \n" +
			"      \n" +
			"      "
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}

func TestConvertToAsciiArtWithShadow(t *testing.T) {
	got := СonvertToAsciiArt("A", _const.ResourcesPath+"shadow.txt")
	want :=
		"         \n" +
			"  _|_|   \n" +
			"_|    _| \n" +
			"_|_|_|_| \n" +
			"_|    _| \n" +
			"_|    _| \n" +
			"         \n" +
			"         "
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}
func TestCreateMap(t *testing.T) {
	filename := _const.ResourcesPath + "standard.txt"
	asciiMap := createMap(filename)
	for char := _const.AsciiSpace; char <= _const.AsciiTilde; char++ {
		if len(asciiMap[char]) == 0 {
			t.Fatalf("\ncreateMap() didn't generate a mapping for rune %v", char)
		}
	}
}
