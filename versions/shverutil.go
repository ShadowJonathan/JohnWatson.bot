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
	for ver, col := range versions {
		if colorarray == col {
			return ver
		}
	}
	// failed, return nil Version
	return Version{0, 0, 0, 0}
}
