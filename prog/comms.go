package prog

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Cute(com []string, m *discordgo.Message) {
	var k bool
	var s int
	if len(com) > 1 && Clearance(1, getguild(m.ChannelID).ID, m.Author.ID)[m.Author.ID] {
		if com[1] == "add" {
			_, ok := J.Data["cute"]
			if !ok {
				J.Data["cute"] = make(map[string]string)
			}
			if len(m.Attachments) > 0 {
				image := m.Attachments[0]
				wot, img, k, s := downloadimg(image.URL)
				if !k {
					handletoobig(m, s)
					return
				}
				if !wot {
					panic(m)
				}
				fn := newfile("cute")
				saveimage("cute", fn, img)
				J.Data["cute"][fn] = ""
				inform(m.ChannelID, "`Added to cute.`")
				go fmt.Println(m.Author.Username + " added cute via url")
			} else if len(com) > 2 {
				image := com[2]
				wot, img, k, s := downloadimg(image)
				if !k {
					handletoobig(m, s)
					return
				}
				if !wot {
					panic(m)
				}
				fn := newfile("cute")
				saveimage("cute", fn, img)
				J.Data["cute"][fn] = image
				inform(m.ChannelID, "`Added to cute.`")
				go fmt.Println(m.Author.Username + " added cute via upload")
			} else {
				inform(m.ChannelID, "`No image provided.`")
			}
			return
		} else if com[1] == "del" {
			if len(com) > 2 {
				_, ok := J.Data["cute"][com[2]]
				if !ok {
					inform(m.ChannelID, "`Image "+com[2]+" does not exist.`")
				} else {
					delete(J.Data["cute"], com[2])
					inform(m.ChannelID, "`Image deleted.`")
					go fmt.Println(m.Author.Username + " removed cute " + com[2])
				}
			} else {
				inform(m.ChannelID, "`No image provided`")
			}
			return
		} else if com[1] == "list" {
			var tot []string
			for d := range J.Data["cute"] {
				tot = append(tot, d)
			}
			inform(m.ChannelID, "`The cute collection:`\n`"+strings.Join(tot, ", ")+"`")
			return
		}
	}
	c, ok := J.Data["cute"]
	if ok {
		var (
			im string
			fb string
		)
		if len(com) < 2 {
			im, fb = MGRI(c)
		} else {
			fb = c[com[1]]
			im = com[1]
		}
		tb := toobig("cute", im)
		if !tb || fb == "" {
			OK, img := getimage("cute", im)
			var last bool
			if !OK {
				last, img, k, s = downloadimg(fb)
				if !k {
					handletoobig(m, s)
					return
				}
				if !last && len(com) < 2 {
					notifyowner("`A collection item from cute has just been invaluated: "+im+", "+fb+"`", m.ChannelID)
					delete(J.Data["cute"], im)
					Cute(com, m)
					return
				} else if !last && len(com) > 1 {
					inform(m.ChannelID, "`Sorry, that cute does not exist!`")
					return
				}
			}
			b := new(bytes.Buffer)
			jpeg.Encode(b, img, &jpeg.Options{Quality: 90})
			d, _ := ioutil.ReadAll(b)
			upload("cute.jpg", d, m.ChannelID)
		} else {
			inform(m.ChannelID, fb)
		}
	} else {
		notifyowner("`Someone tried to use the cute collection, however, this has not been set, please add images`", m.ChannelID)
		inform(m.ChannelID, "`The cute collection doesn't exist!`")
	}
}

func Sona(com []string, m *discordgo.Message) {
	if len(com) > 1 && Clearance(1, getguild(m.ChannelID).ID, m.Author.ID)[m.Author.ID] {
		if com[1] == "add" {
			_, ok := J.Data["sona"]
			if !ok {
				J.Data["sona"] = make(map[string]string)
			}
			if len(com) > 2 {
				user := com[2]
				fmt.Println(user)
				user = parsemention(user)
				fmt.Println(user)
				if !isnumberstring(user) {
					mul := findmem(user)
					if len(mul) < 1 {
						inform(m.ChannelID, "`Can't find user`")
						return
					}
					user = mul[0]
				}

				if len(m.Attachments) < 1 && len(com) < 3 {
					inform(m.ChannelID, "`No image provided.`")
				} else if len(m.Attachments) > 0 {
					image := m.Attachments[0]
					wot, img, k, s := downloadimg(image.URL)
					if !k {
						handletoobig(m, s)
						return
					}
					if !wot {
						panic(m)
					}
					fn := newfile("sona")
					saveimage("sona", fn, img)
					J.Data["sona"][user] = fn
					inform(m.ChannelID, "`Added sona.`")
					go fmt.Println(m.Author.Username + " added sona for " + user)
				} else if len(com) > 2 {
					image := com[3]
					wot, img, k, s := downloadimg(image)
					if !k {
						handletoobig(m, s)
						return
					}
					if !wot {
						panic(m)
					}
					fn := newfile("sona")
					saveimage("sona", fn, img)
					J.Data["sona"][user] = fn
					inform(m.ChannelID, "`Added sona.`")
					go fmt.Println(m.Author.Username + " added sona for " + user)
				}
			} else {
				inform(m.ChannelID, "`No destination user provided`")
			}
			return
		} else if com[1] == "del" {
			if len(com) > 2 {
				user := com[2]
				user = parsemention(user)
				if !isnumberstring(user) {
					mul := findmem(user)
					if len(mul) < 1 {
						inform(m.ChannelID, "`Can't find user`")
						return
					}
					user = mul[0]
				}

				delete(J.Data["sona"], user)
				inform(m.ChannelID, "`Sona deleted.`")
				go fmt.Println(m.Author.Username + " deleted sona for " + user)
			} else {
				inform(m.ChannelID, "`No destination user provided`")
			}
			return
		} else if com[1] == "list" {
			var tot [][2]string
			for u, s := range J.Data["sona"] {
				var l [2]string
				l[0] = u
				l[1] = s
				tot = append(tot, l)
			}
			var T string
			for _, l := range tot {
				T = T + "`\n`" + l[0] + ": " + l[1]
			}
			T = T + "`"
			inform(m.ChannelID, "`The sona collection:"+T)
			return
		}
	}

	sm, ok := J.Data["sona"]
	if ok {
		if len(com) < 2 {
			inform(m.ChannelID, "`No user provided.`")
			return
		}
		user := com[1]
		if user != "rand" && user != "self" {
			if len(user) > 3 {
				user = parsemention(user)
			}
			if !isnumberstring(user) {
				mul := findmem(user)
				if len(mul) < 1 {
					inform(m.ChannelID, "`Can't find user`")
					return
				}
				user = mul[0]
			}

			s, ok := sm[user]
			if ok {
				ye, img := getimage("sona", s+".png")
				if !ye {
					notifyowner("`A collection item from sona has just been invaluated: "+user+" "+s+"`", m.ChannelID)
					delete(J.Data["sona"], user)
					inform(m.ChannelID, "`Oops, i found a sona, but the image linked to it doesnt exist anymore!`")
					return
				}
				b := new(bytes.Buffer)
				png.Encode(b, img)
				d, _ := ioutil.ReadAll(b)
				upload("sona_"+user+".png", d, m.ChannelID)
			} else {
				inform(m.ChannelID, "`No sona found`")
			}
		} else if user == "rand" {
			u, s := MGRI(sm)
			U, _ := J.dg.GuildMember(getguild(m.ChannelID).ID, u)
			ye, img := getimage("sona", s+".png")
			if !ye {
				notifyowner("`A collection item from sona has just been invaluated: "+u+" "+s+"`", m.ChannelID)
				delete(J.Data["sona"], u)
				inform(m.ChannelID, "`Oops, i found a sona, but the image linked to it doesnt exist anymore!`")
				return
			}
			b := new(bytes.Buffer)
			png.Encode(b, img)
			d, _ := ioutil.ReadAll(b)
			if U.Nick != "" {
				uploadmessage("sona_"+u+".png", d, m.ChannelID, "**"+U.Nick+"**'s sona:")
			} else {
				uploadmessage("sona_"+u+".png", d, m.ChannelID, "**"+U.User.Username+"**'s sona:")
			}
		} else if user == "self" {
			if len(com) > 2 || len(m.Attachments) > 0 {
				user = m.Author.ID
				if len(m.Attachments) > 0 {
					image := m.Attachments[0]
					wot, img, k, s := downloadimg(image.URL)
					if !k {
						handletoobig(m, s)
						return
					}
					if !wot {
						panic(m)
					}
					fn := newfile("sona")
					saveimage("sona", fn, img)
					J.Data["sona"][user] = fn
					inform(m.ChannelID, "`Added sona.`")
					go fmt.Println(m.Author.Username + " added sona for " + user)
				} else if len(com) > 2 {
					image := com[2]
					wot, img, k, s := downloadimg(image)
					if !k {
						handletoobig(m, s)
						return
					}
					if !wot {
						panic(m)
					}
					fn := newfile("sona")
					saveimage("sona", fn, img)
					J.Data["sona"][user] = fn
					inform(m.ChannelID, "`Added sona.`")
					go fmt.Println(m.Author.Username + " added sona for " + user)
				}
			} else {
				s, ok := sm[m.Author.ID]
				if ok {
					ye, img := getimage("sona", s+".png")
					if !ye {
						notifyowner("`A collection item from sona has just been invaluated: "+user+" "+s+"`", m.ChannelID)
						delete(J.Data["sona"], user)
						inform(m.ChannelID, "`Oops, i found your sona, but the image linked to it doesnt exist anymore!`")
						return
					}
					b := new(bytes.Buffer)
					png.Encode(b, img)
					d, _ := ioutil.ReadAll(b)
					uploadmessage("sona_"+user+".png", d, m.ChannelID, "`Your sona:`")
				} else {
					inform(m.ChannelID, "`You don't have a sona saved!`")
				}
			}
		}
	} else {
		notifyowner("`Someone tried to use the sona collection, however, this has not been set, please add images to users`", m.ChannelID)
		inform(m.ChannelID, "`The sona collection doesnt exist!`")
	}
}
