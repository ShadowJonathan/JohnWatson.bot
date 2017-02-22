package main

import (
	"JohnWatson.bot/versions"
	"fmt"
	"testing"
)

func TestEncodefile2(t *testing.T) {
	EncodefileBasic("Sherlock.jpg", "Encoded.png")
}

func TestDecodeFile(t *testing.T) {
	ver := DecodeFile("SH.png")
	fmt.Println(ver)
}

func TestEncodefileVersion(t *testing.T) {
	EncodefileVersion("sherlock.jpg", "SH.png", versions.Version{0, 0, 2, 0})
}
