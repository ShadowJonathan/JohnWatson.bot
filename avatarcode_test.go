package main

import (
	"testing"
	"fmt"
)

func TestEncodefile2(t *testing.T) {
	EncodefileBasic("Sherlock.jpg", "Encoded.png")
}

func TestDecodeFile(t *testing.T) {
	fmt.Println(DecodeFile("versions/result.png"))
}