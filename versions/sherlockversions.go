package versions

import "image/color"

type Version struct {
	major        int
	minor        int
	build        int
	experimental int
}

var white = color.RGBA{255, 255, 255, 255}
var black = color.RGBA{0, 0, 0, 255}

var versions = map[Version][25]color.Color{
	{0, 0, 1, 0}: [25]color.Color{
		white, white, white, white, white,
		white, white, white, white, white,
		white, white, white, white, white,
		white, white, white, white, white,
		white, white, white, white, white},
	/*
		00000
		00000
		00000
		00000
		00000
	*/
	{0, 0, 2, 0}: [25]color.Color{
		white, white, white, white, white,
		white, black, white, black, white,
		white, white, white, white, white,
		white, black, white, black, white,
		white, white, white, white, white},
	/*
		00000
		0*0*0
		00000
		0*0*0
		00000
	*/

}
