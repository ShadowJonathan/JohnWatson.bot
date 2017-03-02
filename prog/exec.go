package prog

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var J *Bot

func I(b *Bot, t string) bool {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var err error
	J = b
	if t == "" {
		b, err := ioutil.ReadFile("token")
		if err != nil {
			log.Println(err)
			log.Fatal("Could not find token string nor token file, provide one, please.")
		}
		t = strings.TrimSpace(string(b))
	}

	J.dg, err = discordgo.New("Bot " + t)
	HE(err)
	dg := J.dg

	dg.AddHandler(Chat)
	dg.Open()
	fmt.Println("Opened")
	for !J.Stop && !J.Restart {
		time.Sleep(1 * time.Second)
	}
	dg.Close()
	fmt.Println("Stopping...")
	return J.Restart
}

func Chat(Ses *discordgo.Session, MesC *discordgo.MessageCreate) {
	if len(MesC.Content) > 0 {
		if MesC.Content[0] == '!' {
			ParseCMD(MesC.Message)
			save()
		}
	}
}

func ParseCMD(m *discordgo.Message) {
	M := m.Content[1:]
	op := strings.Split(M, " ")
	ch, _ := J.dg.Channel(m.ChannelID)
	if !ch.IsPrivate {
		if Clearance(0, GetGuildchan(m.ChannelID).ID, m.Author.ID)[m.Author.ID] {
			switch op[0] {
			case "cute":
				Cute(op, m)
			case "fursona":
				Sona(op, m)
			case "sona":
				Sona(op, m)
			case "stop":
				ch, _ := J.dg.Channel(m.ChannelID)
				if !ch.IsPrivate {
					if Clearance(4, getguild(m.ChannelID).ID, m.Author.ID)[m.Author.ID] {
						inform(m.ChannelID, "`Stopping...`")
						J.Stop = true
					}
				} else if m.Author.ID == J.Owner.ID {
					inform(m.ChannelID, "`Stopping...`")
					J.Stop = true
				}
			case "restart":

				ch, _ := J.dg.Channel(m.ChannelID)
				if !ch.IsPrivate {
					if Clearance(3, getguild(m.ChannelID).ID, m.Author.ID)[m.Author.ID] {
						inform(m.ChannelID, "`Restarting...`")
						J.Restart = true
					}
				} else if m.Author.ID == J.Owner.ID {
					inform(m.ChannelID, "`Restarting...`")
					J.Restart = true
				}
			}
		}
	}
}
