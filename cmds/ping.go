package cmds

import (
	"time"

	dg "github.com/bwmarrin/discordgo"
)

type ping struct{}

func (*ping) Info(...string) string {
	return "Usage: <prefix>ping\nsimple command for testing delay of the bot"
}

func (*ping) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	start, _ := m.Timestamp.Parse()
	delay := time.Since(start)

	s.ChannelMessageSendEmbed(m.ChannelID, &dg.MessageEmbed{
		Color: 0xFF00FF,
		Title: "Pong!",
		Author: &dg.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Description: "Ping: " + delay.String(),
	})
}