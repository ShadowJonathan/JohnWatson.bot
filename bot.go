package main

import (
	"encoding/json"
	"io/ioutil"
	"JohnWatson.bot/versions"
)

var version = versions.Version{0,0,1,0}

var John *Bot

func main() {
	data, err := ioutil.ReadFile("Settings.json")
	HE(err)
	sett := &settings{}
	err = json.Unmarshal(data, sett)
	HE(err)
	John = &Bot{
		SherlockAuthlevel: sett.Authlevel,
		Version:           version,
		Owner:             sett.Owner,
	}
}
