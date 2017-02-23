package prog

import (
	"encoding/json"
	"testing"
	"io/ioutil"
)

func TestJson(t *testing.T) {
	data, _ := json.Marshal(&Settings{})
	ioutil.WriteFile("Settings.json", data, 9001)
}
