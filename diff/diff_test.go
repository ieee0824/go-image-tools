package diff

import (
	"image/color"
	"testing"
)

var IsEqualColor = isEqualColor

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
