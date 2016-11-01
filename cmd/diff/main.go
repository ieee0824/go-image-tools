package main

import (
	"image/png"
	"log"
	"os"

	"github.com/ieee0824/go-image-tools/diff"
)

func main() {
	file0, _ := os.Open("test_0.png")
	file1, _ := os.Open("test_1.png")

	imgA, _ := png.Decode(file0)
	imgB, _ := png.Decode(file1)

	result, err := diff.Diff(imgA, imgB)
	if err != nil {
		log.Fatalln(err)
	}
	out, _ := os.Create("result.png")

	png.Encode(out, result)

}
