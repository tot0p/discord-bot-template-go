package discord

import (
	"discord/command"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix      = "/" // default prefix
	listcommand = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string){
		"ping": command.Ping,
	} // default command
)

// AddCommand add command to listcommand
func AddCommand(name string, f func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)) {
	listcommand[name] = f
}

// DeleteCommand delete command from listcommand
func DeleteCommand(name string) {
	delete(listcommand, name)
}

func SetPrefix(p string) {
	prefix = p
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// split message to get  keyword and args
	temp := strings.Split(m.Content, " ")
	keyword := temp[0]
	var args = []string{}
	if len(temp) > 1 {
		args = temp[1:]
	}
	// check if message is command
	if strings.Contains(keyword, prefix) {
		// call command if exist
		if f, ok := listcommand[strings.Replace(keyword, prefix, "", 1)]; ok {
			f(s, m, args)
		}
	}
}
