package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	data := []int{100, 33, 73, 64}

	w, h := len(data)*60+10, 100
	r := image.Rect(0, 0, w, h)
	img := image.NewRGBA(r)

	// for y := 0; y < h; y++ {
	// 	for x := 0; x < w; x++ {
	// 		img.Set(x, y, color.RGBA{
	// 			R: uint8((x + y) & 255),
	// 			G: uint8((x + y) << 1 & 255),
	// 			B: uint8((x + y) << 2 & 255),
	// 			A: 255,
	// 		})
	// 	}
	// }

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	for i, dp := range data {
		for x := i*60 + 10; x < (i+1)*60; x++ {
			for y := 100; y >= (100 - dp); y-- {
				img.Set(x, y, color.RGBA{180, 180, 250, 255})
			}
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
