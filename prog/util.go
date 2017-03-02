package prog

import (
	"bytes"
	ran "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HE(err error) {
	if err != nil {
		panic(err)
	}
}

func MGRI(m map[string]string) (string, string) {
	index := rand.Intn(len(m))
	for k, v := range m {
		if index == 0 {
			return k, v
		}
		index--
	}

	return "ERROR", ""
}

func getimage(source, im string) (bool, image.Image) {
	f, err := os.Open(source + "/" + im)
	defer f.Close()
	if err != nil {
		return false, nil
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return false, nil
	} else {
		return true, img
	}
}

func toobig(source, im string) bool {
	f, err := os.Open(source + "/" + im + ".png")
	defer f.Close()
	if err != nil {
		return false
	}
	fi, err := f.Stat()
	HE(err)
	ye := fi.Size() > 1000000
	fmt.Println(fi.Size())
	return ye
}

var DLlimit = 2000000

func downloadimg(url string) (bool, image.Image, bool, int) {
	res, err := download(url)
	if err != nil {
		return false, nil, false, 0
	}
	b := res.Body
	defer b.Close()
	img, _, err := image.Decode(b)
	if err != nil {
		return false, nil, false, 0
	} else {
		ch, _ := download(url)
		body, err := ioutil.ReadAll(ch.Body)
		defer ch.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		l := len(body)
		fmt.Println(l)
		if l > DLlimit {
			return true, img, false, l
		} else {
			return true, img, true, l
		}
	}
}

func handletoobig(m *discordgo.Message, l int) {
	if l > 1 {
		inform(m.ChannelID, "`The image that you gave me is too big!`\n`It's "+Getdatacount(int64(l))+", but it has to be lower than 2 MB!`")
	} else {
		inform(m.ChannelID, "`That link doesn't point to a valid download, please check the link, and if it has to be an image, please link the image directly.`")
	}
}

func Getdatacount(i int64) string {
	s := strconv.FormatInt(i, 10)
	if i > 1000000 {
		t := s[len(s)-6:]
		h := s[:len(s)-6]
		if t[1] == '0' {
			return h + "." + t[:1] + " MB"
		} else {
			return h + "." + t[:2] + " MB"
		}
	} else {
		return s + " KB"
	}
	return s
}

func download(url string) (*http.Response, error) {
	return http.Get(url)
}

func upload(n string, d []byte, ch string) {
	uploadmessage(n, d, ch, "")
}

func uploadmessage(n string, d []byte, ch string, mes string) {
	b := new(bytes.Buffer)
	b.Write(d)
	J.dg.ChannelFileSendWithMessage(ch, mes, n, b)
}

func inform(ch string, mes string) {
	_, err := J.dg.ChannelMessageSend(ch, mes)
	if err != nil {
		fmt.Println(err)
	}
}

func notifyowner(mes string, ch string) {
	g := GetGuildchan(ch)
	owner, ok := J.Data["owners"][g.ID]
	if !ok {
		return
	}
	opch, err := J.dg.UserChannelCreate(owner)
	if err != nil {
		return
	}
	inform(opch.ID, mes)
}

func notifymods(mes string, ch string) {
	g := GetGuildchan(ch)
	mods, ok := J.Data["MODS "+g.ID]
	if !ok {
		return
	}
	for mod := range mods {
		opch, err := J.dg.UserChannelCreate(mod)
		if err != nil {
			continue
		}
		inform(opch.ID, mes)
	}
}

func GetGuildchan(ch string) *discordgo.Guild {
	for _, g := range J.dg.State.Guilds {
		for _, c := range g.Channels {
			if ch == c.ID {
				return g
			}
		}
	}
	return nil
}

func isnumberstring(s string) bool {
	var check = true
	for _, l := range s {
		if !strings.ContainsAny(string(l), "0 1 2 3 4 5 6 7 8 9") {
			check = false
		}
	}
	return check
}

func parsemention(s string) string {
	if s[:3] == "<@!" && s[len(s)-1] == '>' {
		s = s[3 : len(s)-1]
	}
	if s[:2] == "<@" && s[len(s)-1] == '>' {
		s = s[2 : len(s)-1]
	}
	return s
}

func findmem(name string) []string {
	var all []string
	for _, guild := range J.dg.State.Guilds {
		for _, m := range guild.Members {
			if (strings.ToLower(name) == strings.ToLower(m.User.Username)) || (strings.ToLower(name) == strings.ToLower(m.Nick)) || (name == m.User.ID) {
				all = append(all, m.User.ID)
			}
		}
	}
	return all
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := ran.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func newfile(source string) string {
	fn, _ := GenerateRandomString(8)
	_, err := ioutil.ReadFile(source + "/" + fn)
	if err == nil {
		return newfile(source)
	} else {
		return fn
	}
}

func saveimage(source, name string, img image.Image) {
	os.MkdirAll(source, 0777)
	f, err := os.Create(source + "/" + name + ".png")
	defer f.Close()
	HE(err)
	png.Encode(f, img)
}

func getguild(ch string) *discordgo.Guild {
	c, err := J.dg.Channel(ch)
	if err != nil {
		return nil
	} else {
		g, err := J.dg.Guild(c.GuildID)
		if err != nil {
			return nil
		} else {
			return g
		}
	}

}

func Clearance(i int, g string, people ...string) map[string]bool {
	cc := J.Data["clear"]
	var ret = make(map[string]bool)
	var c int
	for _, p := range people {
		cs, ok := cc[p]
		if !ok {
			c = -1
		} else {
			c64, _ := strconv.ParseInt(cs, 10, 0)
			c = int(c64)
		}
		r := Clroles(p, g)
		if r > c {
			c = r
		}
		if i <= c {
			ret[p] = true
		} else {
			ret[p] = false
		}

	}
	return ret
}

func Clroles(p string, g string) int {
	m, err := J.dg.GuildMember(g, p)
	if err != nil {
		return -1
	}
	r := m.Roles
	var max int = -1
	for _, role := range r {
		c, ok := J.Data["clear"][role]
		if ok {
			i64, err := strconv.ParseInt(c, 10, 0)
			if err != nil {
				continue
			}
			i := int(i64)
			if i > max {
				max = i
			}
		} else {
			continue
		}
	}
	return max
}

func save() {
	b, err := ioutil.ReadFile("Settings.json")
	if err != nil {
		panic(err)
	}
	var s = &Settings{}
	err = json.Unmarshal(b, s)
	if err != nil {
		fmt.Println(err)
	}
	if J.Owner != nil {
		s.Owner = J.Owner
	}
	s.Authlevel = J.SherlockAuthlevel
	if J.Data != nil {
		s.Data = J.Data
	}
	if J.Sherlocks != nil {
		s.Sherlocks = J.Sherlocks
	}
	data, _ := json.Marshal(s)
	ioutil.WriteFile("Settings.json", data, 0666)
}
