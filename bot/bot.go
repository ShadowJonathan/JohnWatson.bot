package main

import (
	"JohnWatson.bot/versions"
	"encoding/json"
	"io/ioutil"
	"JohnWatson.bot/prog"
)

var version = versions.Version{0, 0, 1, 0}

var John *prog.Bot

func main() {
	data, err := ioutil.ReadFile("Settings.json")
	prog.HE(err)
	sett := &prog.Settings{}
	err = json.Unmarshal(data, sett)
	prog.HE(err)
	John = &prog.Bot{
		SherlockAuthlevel: sett.Authlevel,
		Version:           version,
		Owner:             sett.Owner,
	}
}
