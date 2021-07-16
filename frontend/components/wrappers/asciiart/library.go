package asciiart

import (
	"embed"
	"fmt"
	"math"
	"strings"
)

//go:embed *
var media embed.FS

type Library string

const (
	KEYS        Library = "arrowkeys"
	TITLE       Library = "title"
	ENTERKEY    Library = "enterkey"
	SPACEBARKEY Library = "spacebarkey"
)

type AsciiArt struct {
	title  string
	width  int
	height int
}

func (t AsciiArt) String() string {
	return t.title
}

func (t AsciiArt) Width() int {
	return t.width

}

func (t AsciiArt) Height() int {
	return t.height
}

// New returns an ascii that fits in the given dimensions
func New(library Library, width, height int) (AsciiArt, error) {
	if width <= 0 || height <= 0 {
		return AsciiArt{}, fmt.Errorf("invalid dimensions %d, %d", width, height)
	}
	title, w, h := search(library, width, height)
	return AsciiArt{
		title:  title,
		width:  w,
		height: h,
	}, nil
}

// search returns a title that fits in the given dimensions.
// returns an empty title if no text meets width and height conditions.
func search(library Library, width, height int) (string, int, int) {
	dir, err := media.ReadDir(string(library))
	if err != nil {
		panic("cannot find ascii")
	}
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		bytes_, err := media.ReadFile(fmt.Sprintf("%s/%s", library, entry.Name()))
		if err != nil {
			panic(fmt.Sprintf("cannot read file %s", entry.Name()))
		}
		text := string(bytes_)
		w, h := getDimensions(text)
		if w <= width && h <= height {
			return text, w, h
		}
	}
	return "", 0, 0
}

// getDimensions return text's width and height
func getDimensions(text string) (int, int) {
	lines := strings.Split(text, "\n")
	height := len(lines)
	width := 0
	for _, line := range lines {
		width = int(math.Max(float64(len(line)), float64(width)))
	}
	return width, height
}
