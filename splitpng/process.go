package main

import (
	"fmt"
	"image"
	"log"
)

func slice(i image.Image) []image.Image {
	results := make([]image.Image, 0)
	px := findAvatarsPixel(i)
	top := 0
	bottom := 0
	nrgba, ok := i.(*image.NRGBA)
	if !ok {
		log.Fatalf("type of i is %T not image.RGBA", i)
	}
	for j := i.Bounds().Min.Y; j <= i.Bounds().Max.Y; j++ {
		if lum(i.At(px, j)) < 255.0 {
			// found an avatar
			if top == 0 && bottom == 0 {
				// special case: the first vertical avatar does not mark a comment end. skip it and cut at the next avatar
				top = j - 5 // padding
				j += 50
				continue
			}

			// first vertical pixel of a subsequent avatar
			bottom = j - 5
			i2 := nrgba.SubImage(
				image.Rectangle{
					image.Point{i.Bounds().Min.X, top},    // top left above avatar, with 5px padding
					image.Point{i.Bounds().Max.X, bottom}, // bottom left above next comments avatar, with 5px padding
				},
			)
			fmt.Printf("found comment of height %v\n", i2.Bounds().Max.Y-i2.Bounds().Min.Y)
			results = append(results, i2)
			top = j - 5 // padding above next comment
			j += 50     // skip over the 40px avatar
			// fix final comment off-by-one:
			if j > i.Bounds().Max.Y && j < i.Bounds().Max.Y+50 {
				j = i.Bounds().Max.Y
			}

		}
		// luminance was 1; whitespace in avatar column, keep going down
	}
	return results
}

// given an image of youtube comments pick a pixel in the middle of the avatar column
func findAvatarsPixel(i image.Image) int {
	result := sum(i)
	normalized := normalize(result)
	l, r := 0, 0
	for i, Yi := range normalized {
		// if initial whitespace: continue
		if l == 0 && r == 0 && Yi == 1.0 {
			continue
		}
		// if first avatar pixel: set left, continue
		if l == 0 {
			l = i
			continue
		}
		// mid or last avatar pixel
		r = i
		if normalized[i+1] == 1.0 {
			break
		}
	}
	return l + ((r - l) / 2)
}
