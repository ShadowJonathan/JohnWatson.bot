package prog

import (
	"time"

	"../versions"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	ID                string
	SelfUser          *discordgo.User
	SherlockAuthlevel int // 0: all data, 1: only chat, 2: only changes, 3: none
	Sherlocks         []*Sherlock
	Version           versions.Version //major, minor, build, experimental
	Owner             *discordgo.User
	dg                *discordgo.Session
	Data              map[string]map[string]string
	Stop              bool
	Restart           bool
}

type Sherlock struct {
	ID           string
	Active       bool
	Lastseen     time.Time
	Undercover   bool // is this sherlock a userbot, yes/no
	Autodetected bool // is this bot seen by the avatar recognition, yes/no
	Owner        *discordgo.User
	Protocol     string
	Version      versions.Version
}

type Settings struct {
	Token     string                       `json:"token"`
	Authlevel int                          `json:"Auth"`
	Owner     *discordgo.User              `json:",omitempty"`
	Data      map[string]map[string]string `json:"data,omitempty"`
	Sherlocks []*Sherlock                  `json:",omitempty"`
}
