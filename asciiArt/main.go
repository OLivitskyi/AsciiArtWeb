package asciiArt

import (
	_const "ASCIIArt/asciiArt/constants"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if at least one argument is present
	if len(os.Args) <= _const.ArgInputText {
		fmt.Println("Error: not enough arguments")
		fmt.Println("Usage: go run main.go [TEXT_TO_PRINT] [OPTIONAL_ASCII_ART_STYLE]")
		fmt.Println("Available ASCII Art styles: standard, shadow, thinkertoy")
		return
	}

	text := os.Args[_const.ArgInputText]
	bannerFilename := _const.DefaultBannerName + ".txt"

	if len(os.Args) > _const.ArgBannerName {
		requestedBanner := os.Args[_const.ArgBannerName]
		availableBanners := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}
		if availableBanners[requestedBanner] {
			bannerFilename = requestedBanner + ".txt"
		} else {
			fmt.Println("The Banner", requestedBanner, "is not available. Using", _const.DefaultBannerName+".txt", "as a default ASCII Art.")
		}
	}

	asciiArt := СonvertToAsciiArt(text, _const.ResourcesPath+bannerFilename)
	fmt.Println(asciiArt)
}

func СonvertToAsciiArt(text, bannerFilename string) string {
	text = strings.Replace(text, "\\n", "\n", -1)
	asciiMap := createMap(bannerFilename)
	lines := strings.Split(text, "\n")
	var artLines []string
	for _, line := range lines {
		artLines = append(artLines, printArt([]string{line}, asciiMap))
	}
	return strings.Join(artLines, "\n")
}

func createMap(filename string) map[rune][]string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	bannerLines := strings.Split(string(fileContent), "\n")
	asciiMap := make(map[rune][]string, _const.AsciiSymbolsCount)

	totalLines := len(bannerLines)

	linesPerSymbol := _const.WithoutSpacingSeparator
	if (totalLines / _const.AsciiSymbolsCount) == _const.WithSpacingSeparator {
		linesPerSymbol = _const.WithSpacingSeparator
	}

	for char := _const.AsciiSpace; char <= _const.AsciiTilde; char++ {
		startLine := _const.AsciiArtStartOffset + (int(char)-int(_const.AsciiSpace))*linesPerSymbol
		endLine := startLine + _const.RowCountPerChar
		fullArt := bannerLines[startLine:endLine]
		asciiMap[char] = fullArt
	}

	return asciiMap
}

func printArt(lines []string, asciiMap map[rune][]string) string {
	var result string
	for lineIndex, word := range lines {
		if word == "" {
			continue
		}
		for row := 0; row < _const.RowCountPerChar; row++ {
			for _, ch := range word {
				result += asciiMap[ch][row]
			}
			if lineIndex != len(lines)-1 || row != _const.RowCountPerChar-1 {
				result += "\n"
			}
		}
	}
	return result
}
