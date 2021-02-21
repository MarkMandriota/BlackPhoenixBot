package cmds

import (
	"strings"

	dg "github.com/bwmarrin/discordgo"
)

type help struct{}

func (*help) Info(a ...string) string {
	switch len(a) {
	case 0:
		return `Usage: <prefix>help <?command>
				Show info about commands`
	}
	return Map[a[0]].Info(a[1:]...)
}

func (*help) Perm(a ...string) int64 {
	return 0
}

func (*help) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	msg := &dg.MessageEmbed{
		Color: 0xFF00FF,
		Author: &dg.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
	}

	if len(a) > 1 {
		msg.Title = strings.Join(a[1:], " ")
		msg.Description = Map[a[1]].Info(a[2:]...)
	} else {
		msg.Title = `Help`
		for k, v := range Map {
			msg.Fields = append(msg.Fields, &dg.MessageEmbedField{
				Name:   k,
				Value:  v.Info(),
				Inline: true,
			})
		}
	}

	s.ChannelMessageSendEmbed(m.ChannelID, msg)
}
