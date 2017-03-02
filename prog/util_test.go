package prog

import (
	"testing"
	"image"
)

func TestSaveImage(t *testing.T) {
	var img image.Image
	saveimage("cute", "test", img)
}
