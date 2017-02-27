package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
)

func Encodefile(file, dest string) {
	data, err := os.Open(file)
	HE(err)
	img, _, err := image.Decode(data)
	HE(err)
	img = Encode(img)
	out, err := os.Create(dest)
	png.Encode(out, img)
}

func Encode(img image.Image) image.Image {
	e := &encodedimage{}
	e.Image = img
	return e
}

type encodedimage struct {
	image.Image
}

func (m *encodedimage) At(x, y int) color.Color {
	// "Changed" part: custom colors for specific coordinates:
	if x <= 5 && y <= 5 {
		return color.RGBA{255, 255, 255, 255}
	}
	// "Unchanged" part: the colors of the original image:
	return m.Image.At(x, y)
}
