package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	_, err := s.ChannelMessageSend(m.ChannelID, "pong")
	if err != nil {
		fmt.Println(err)
	}
}
