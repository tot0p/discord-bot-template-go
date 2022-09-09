package discord

import (
	"discord/command"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix      = "!"
	listcommand = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string){
		"ping": command.Ping,
	}
)

func AddCommand(name string, f func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {
	listcommand[name] = f
}

func SetPrefix(p string) {
	prefix = p
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	temp := strings.Split(m.Content, " ")
	keyword := temp[0]
	var args = []string{}
	if len(temp) > 1 {
		args = temp[1:]
	}
	if strings.Contains(keyword, prefix) {
		if f, ok := listcommand[strings.Replace(keyword, prefix, "", 1)]; ok {
			f(s, m, args)
		}
	}
}
