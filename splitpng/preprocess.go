package main

import (
	"image"
	"image/color"
)

// sums vertically, returning summed pixel luminance row in range 0..1
func sum(i image.Image) []int64 {
	ll := make([]int64, i.Bounds().Max.X)
	for x := 0; x < i.Bounds().Max.X; x++ { // column
		for y := 0; y < i.Bounds().Max.Y; y++ { // row
			ll[x] += lum(i.At(x, y))
		}
	}
	return ll
}

func lum(c color.Color) int64 {
	r, g, b, _ := c.RGBA()
	y, _, _ := color.RGBToYCbCr(uint8(r), uint8(g), uint8(b))
	return int64(y)
}

func max(ff []int64) int64 {
	max := int64(0)
	for _, f := range ff {
		if f > max {
			max = f
		}
	}
	return max
}

func min(ff []int64) int64 {
	min := max(ff)
	for _, f := range ff {
		if f < min {
			min = f
		}
	}
	return min
}

// scale min..Y..max -> 0..1
func scale(Y, min, max int64) float64 {
	rng := max - min
	deltaTop := max - Y
	dy := float64(deltaTop) / float64(rng)
	return 1 - dy
}

func normalize(result []int64) []float64 {
	maxY := max(result)
	minY := min(result)
	out := make([]float64, len(result))
	for i, Y := range result {
		out[i] = scale(Y, minY, maxY)
	}
	return out
}
