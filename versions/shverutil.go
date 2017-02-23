package versions

import (
	"image"
	"image/color"
)

func GetSHVersion(img image.Image) Version {
	colorarray := [25]color.Color{}
	var ai int
	ai = 0
	for i := img.Bounds().Dy() - 5; i < img.Bounds().Dy(); i++ {
		for I := img.Bounds().Dy() - 5; I < img.Bounds().Dy(); I++ {
			colorarray[ai] = img.At(I, i)
			ai++
		}
	}
	fail := Version{0, 0, 0, 0}
	header := colorarray[:4]
	Major := colorarray[4:9]
	Minor := colorarray[9:14]
	Build := colorarray[14:19]
	Exper := colorarray[19:24]
	for _, c := range header {
		if c != white {
			return fail
		}
	}
	return Version{
		convert(Major),
		convert(Minor),
		convert(Build),
		convert(Exper),
	}
}

func convert(c []color.Color) int {
	if len(c) > 5 || len(c) < 5 {
		panic(c)
	}
	var w int = 0
	if c[4] == black {
		w = w+1
	}
	if c[3] == black {
		w = w+2
	}
	if c[2] == black {
		w = w+4
	}
	if c[1] == black {
		w = w+8
	}
	return w
}