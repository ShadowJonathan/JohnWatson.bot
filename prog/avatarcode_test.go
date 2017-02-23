package prog

import (
	"JohnWatson.bot/versions"
	"fmt"
	"testing"
	"os"
	"image"
	"image/png"
)

const this = "../prog/"

func TestEncodefile2(t *testing.T) {
	EncodefileBasic(this+"sherlock.jpg", this+"Encoded.png")
}

func TestEncodefileVersion(t *testing.T) {
	EncodefileVersion("sherlock.jpg", "SH.png", versions.Version{0, 0, 2, 0})
}

func TestDecodeFile(t *testing.T) {
	ver := DecodeFile("SH.png")
	fmt.Println(ver)
}

func TestPushVersion(t *testing.T) {
	data, err := os.Open(this + "sherlock.jpg")
	HE(err)
	img, _, err := image.Decode(data)
	HE(err)
	img = EncodeBasic(img)
	img = EncodeVersion(img, versions.Version{0, 0, 2, 0})
	out, err := os.Create("Encoded.png")
	png.Encode(out, img)
}

var localimage = "http://localhost:9001/SH.png"

func TestDecodeUrl(t *testing.T) {
	fmt.Println(DecodeUrl(localimage))
}