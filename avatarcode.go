package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
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
	if x <= 5 && y <= 5 {
		return color.RGBA{255, 255, 255, 255}
	}
	return m.Image.At(x, y)
}
