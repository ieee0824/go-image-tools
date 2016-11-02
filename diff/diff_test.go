package diff

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
	"testing"
)

var testPath = fmt.Sprintf("%s/src/github.com/ieee0824/go-image-tools/test", os.Getenv("GOPATH"))
var inPath = fmt.Sprintf("%s/in", testPath)
var outPath = fmt.Sprintf("%s/out", testPath)

var IsEqualColor = isEqualColor
var ColorDiff = colorDiff

var (
	RED = color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0x00,
		A: 0xff,
	}
	GREEN = color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	}
	BLUE = color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0xff,
		A: 0xff,
	}
	Transparency = color.RGBA{
		R: 0xff,
		G: 0x00,
		B: 0x00,
		A: 0x00,
	}
)

func TestIsEqualColor(t *testing.T) {
	var actual bool
	actual = IsEqualColor(RED, RED)
	if !actual {
		t.Errorf("got %v\nwant %v\n", actual, true)
	}

	actual = IsEqualColor(RED, GREEN)
	if actual {
		t.Errorf("got %v\nwant %v\n", actual, false)
	}
	actual = IsEqualColor(RED, Transparency)
	if actual {
		t.Errorf("got %v\nwant %v\n", actual, true)
	}
}

func BenchmarkIsEqualColor(b *testing.B) {
	cn := 3
	for i := 0; i < b.N; i++ {
		if i%cn == 0 {
			IsEqualColor(RED, RED)
		} else if i%cn == 1 {
			IsEqualColor(RED, GREEN)
		} else if i%cn == 2 {
			IsEqualColor(RED, BLUE)
		} else if i%cn == 3 {
			IsEqualColor(RED, Transparency)
		}
	}
}

func TestColorDiff(t *testing.T) {
	var result float64

	result = ColorDiff(RED, RED)
	if result != 2.0 {
		t.Errorf("got %v, want %v\n", result, 2.0)
	}

	result = ColorDiff(RED, GREEN)
	if result != 3.0 {
		t.Errorf("got %v, want %v\n", result, 3.0)
	}

	result = ColorDiff(RED, BLUE)
	if result != 3.0 {
		t.Errorf("got %v, want %v\n", result, 3.0)
	}

	result = ColorDiff(RED, Transparency)
	if result != 3.0 {
		t.Errorf("got %v, want %v\n", result, 3.0)
	}
}

func BenchmarkColorDiff(b *testing.B) {
	cn := 3
	for i := 0; i < b.N; i++ {
		if i%cn == 0 {
			ColorDiff(RED, RED)
		} else if i%cn == 1 {
			ColorDiff(RED, GREEN)
		} else if i%cn == 2 {
			ColorDiff(RED, BLUE)
		} else if i%cn == 3 {
			ColorDiff(RED, Transparency)
		}
	}
}

func BenchmarkDiff(b *testing.B) {
	png0f, _ := os.Open(inPath + "/test_0.png")
	png1f, _ := os.Open(inPath + "/test_1.png")

	png0, _ := png.Decode(png0f)
	png1, _ := png.Decode(png1f)

	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			Diff(png0, png1)
		} else if i%3 == 1 {
			Diff(png0, png0)
		} else {
			Diff(png1, png1)
		}
	}
}

func BenchmarkIsEqual(b *testing.B) {
	png0f, _ := os.Open(inPath + "/test_0.png")
	png1f, _ := os.Open(inPath + "/test_1.png")

	png0, _ := png.Decode(png0f)
	png1, _ := png.Decode(png1f)

	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			IsEqual(png0, png1)
		} else if i%3 == 1 {
			IsEqual(png0, png0)
		} else {
			IsEqual(png1, png1)
		}
	}
}

func BenchmarkDifferenceRatioPixel(b *testing.B) {
	png0f, _ := os.Open(inPath + "/test_0.png")
	png1f, _ := os.Open(inPath + "/test_1.png")

	png0, _ := png.Decode(png0f)
	png1, _ := png.Decode(png1f)

	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			DifferenceRatioPixel(png0, png1)
		} else if i%3 == 1 {
			DifferenceRatioPixel(png0, png0)
		} else {
			DifferenceRatioPixel(png1, png1)
		}
	}
}

func BenchmarkDifferenceRatioColor(b *testing.B) {
	png0f, _ := os.Open(inPath + "/test_0.png")
	png1f, _ := os.Open(inPath + "/test_1.png")

	png0, _ := png.Decode(png0f)
	png1, _ := png.Decode(png1f)

	for i := 0; i < b.N; i++ {
		if i%3 == 0 {
			DifferenceRatioColor(png0, png1)
		} else if i%3 == 1 {
			DifferenceRatioColor(png0, png0)
		} else {
			DifferenceRatioColor(png1, png1)
		}
	}
}
