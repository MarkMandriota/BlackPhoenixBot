package cmds

import (
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

type send struct{}

func (*send) Info(...string) string {
	return `Usage: <prefix><message>
			Send message`
}

func (*send) Perm(a ...string) int64 {
	return 0
}

func (*send) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	msg := &dg.MessageEmbed{
		Color: 0xFF00FF,
		Author: &dg.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Description: strings.Join(a[1:], " "),
	}

	s.ChannelMessageSendEmbed(m.ChannelID, msg)
}
