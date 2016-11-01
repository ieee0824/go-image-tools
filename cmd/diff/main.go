package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/ieee0824/go-image-tools/diff"
)

func main() {
	file0, _ := os.Open("test/test_0.jpg")
	file1, _ := os.Open("test/test_1.jpg")

	imgA, _ := jpeg.Decode(file0)
	_ = imgA
	imgB, _ := jpeg.Decode(file1)
	_ = imgB

	r, _ := diff.DifferenceRatioColor(imgA, imgB)
	d, _ := diff.Diff(imgA, imgB)

	out, _ := os.Create("result.png")
	png.Encode(out, d)

	fmt.Println(*r)
}
