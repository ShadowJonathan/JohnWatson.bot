package main

import (
	"../prog"
	"../versions"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"strconv"
)

var version = versions.Version{0, 0, 1, 0}

var John *prog.Bot

func main() {
	data, err := ioutil.ReadFile("Settings.json")
	sett := &prog.Settings{}
	if err != nil {
		sett.Owner = &discordgo.User{}
		sett.Authlevel = 3
		sett.Data = make(map[string]map[string]string)
		fmt.Println(err)
	} else {
		err = json.Unmarshal(data, sett)
		prog.HE(err)
	}
	John = &prog.Bot{
		SherlockAuthlevel: sett.Authlevel,
		Owner:             sett.Owner,
		Version:           versions.Version{0, 0, 1, 0},
		Data:              sett.Data,
		Stop:              false,
		Restart:           false,
	}
	restart := prog.I(John, sett.Token)
	ioutil.WriteFile("../retcmd.botboot", compilebotboot(restart), 0777)
}

func compilebotboot(upgrade bool) []byte {
	return []byte(strconv.FormatBool(upgrade))
}

