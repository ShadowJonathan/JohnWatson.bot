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
	for ver, col := range Vers {
		var ye bool
		ye = true
		for i, c := range col {
			if compcolor(colorarray[i], c) {
				continue
			} else {
				ye = false
			}
		}
		if ye {
			return ver
		}
	}
	// failed, return nil Version
	return Version{0, 0, 0, 0}
}

func compcolor(col1, col2 color.Color) bool {
	r1, b1, g1, a1 := col1.RGBA()
	r2, b2, g2, a2 := col2.RGBA()
	if r1 == r2 && b1 == b2 && g1 == g2 && a1 == a2 {
		return true
	} else {
		return false
	}
}
