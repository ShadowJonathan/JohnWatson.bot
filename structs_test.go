package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestJson(t *testing.T) {
	data, _ := json.Marshal(&settings{})
	ioutil.WriteFile("Settings.json", data, 9001)
}
