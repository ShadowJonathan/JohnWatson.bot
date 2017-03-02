package prog

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestJson(t *testing.T) {
	var s = &Settings{}
	s.Data = make(map[string]map[string]string)
	s.Data["clear"] = make(map[string]string)
	s.Data["clear"]["132583718291243008"] = "4"
	data, _ := json.Marshal(s)
	ioutil.WriteFile("Settings.json", data, 0666)
}
