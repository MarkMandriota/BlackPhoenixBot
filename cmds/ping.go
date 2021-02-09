package cmds

import (
	"time"

	dg "github.com/bwmarrin/discordgo"
)

type ping struct{}

func (p *ping) Info() string {
	return "simple command for testing delay of the bot"
}

func (p *ping) Exec(s *dg.Session, m *dg.MessageCreate) {
	end, _ := m.Timestamp.Parse()
	delay := time.Since(end)

	s.ChannelMessageSendEmbed(m.ChannelID, &dg.MessageEmbed{
		Color:       0xFF00FF,
		Title:       "Pong!",
		Description: "Ping: " + delay.String(),
	})
}
