package diff

import (
	"errors"
	"image"
	"image/color"
)

func isEqualColor(cA, cB color.Color) bool {
	ra, ga, ba, aa := cA.RGBA()
	rb, gb, bb, ab := cB.RGBA()

	return ra == rb && ga == gb && ba == bb && aa == ab
}

func Diff(imgA, imgB image.Image) (*image.RGBA, error) {
	rect := imgA.Bounds()
	ret := image.NewRGBA(rect)
	if !rect.Eq(imgB.Bounds()) {
		return nil, errors.New("not equal emage size")
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if !isEqualColor(imgA.At(x, y), imgB.At(x, y)) {
				ret.Set(x, y, imgA.At(x, y))
			}
		}
	}
	return ret, nil
}

func IsEqual(imgA, imgB image.Image) bool {
	rect := imgA.Bounds()

	if !rect.Eq(imgB.Bounds()) {
		return false
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if !isEqualColor(imgA.At(x, y), imgB.At(x, y)) {
				return false
			}
		}
	}
	return true
}
