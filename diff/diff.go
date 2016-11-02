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

func colorDiff(cA, cB color.Color) float64 {
	ra, ga, ba, aa := cA.RGBA()
	rb, gb, bb, ab := cB.RGBA()

	r := 1 - ratio(ra, rb)
	g := 1 - ratio(ga, gb)
	b := 1 - ratio(ba, bb)
	a := 1 - ratio(aa, ab)

	return (r + g + b + a)
}

func ratio(a, b uint32) float64 {
	if a > b {
		if b == 0 {
			return 0
		}
		return float64(b) / float64(a)
	} else if a == 0 {
		return 0
	}
	return float64(a) / float64(b)
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

func DifferenceRatioPixel(imgA, imgB image.Image) (*float64, error) {
	var result float64
	rect := imgA.Bounds()
	counter := 0

	if !rect.Eq(imgB.Bounds()) {
		return nil, errors.New("not equal emage size")
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if !isEqualColor(imgA.At(x, y), imgB.At(x, y)) {
				counter++
			}
		}
	}

	result = float64(counter) / float64((rect.Max.X-rect.Min.X)*(rect.Max.Y-rect.Min.Y))
	return &result, nil
}

func DifferenceRatioColor(imgA, imgB image.Image) (*float64, error) {
	var result float64
	rect := imgA.Bounds()

	if !rect.Eq(imgB.Bounds()) {
		return nil, errors.New("not equal emage size")
	}
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			if !isEqualColor(imgA.At(x, y), imgB.At(x, y)) {
				result += colorDiff(imgA.At(x, y), imgB.At(x, y))
			}
		}
	}
	result /= float64((rect.Max.X - rect.Min.X) * (rect.Max.Y - rect.Min.Y))
	result /= 4
	return &result, nil
}
