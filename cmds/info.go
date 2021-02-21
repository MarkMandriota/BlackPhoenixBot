package cmds

import (
	"fmt"

	dg "github.com/bwmarrin/discordgo"
)

type info struct{}

func (*info) Info(a ...string) string {
	return `Usage: <prefix>info <?ping/id>
			send info about user`
}

func (*info) Perm(a ...string) int64 {
	return 0
}

func (*info) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	for _, arg := range a[1:] {
		if user, err := s.User(arg); err == nil {
			m.Mentions = append(m.Mentions, user)
		}
	}

	if len(m.Mentions) == 0 {
		m.Mentions = append(m.Mentions, m.Author)
	}

	for _, user := range m.Mentions {
		member, _ := s.GuildMember(m.GuildID, user.ID)
		const inline = true

		msg := &dg.MessageEmbed{
			Color:       0xFF00FF,
			Title:       `Info`,
			Description: `Info about user ` + user.String(),
			Thumbnail: &dg.MessageEmbedThumbnail{
				URL:      user.AvatarURL(""),
				ProxyURL: user.Avatar,
			},
			Fields: []*dg.MessageEmbedField{
				{Name: `Username`, Value: user.Username, Inline: inline},
				{Name: `Discriminator`, Value: user.Discriminator, Inline: inline},
				{Name: `ID`, Value: check(user.ID), Inline: inline},
				{Name: `Bot`, Value: fmt.Sprint(user.Bot), Inline: inline},
				{Name: `Verified`, Value: fmt.Sprint(user.Verified), Inline: inline},
				{Name: `MFAEnabled`, Value: fmt.Sprint(user.MFAEnabled), Inline: inline},
				{Name: `Flags`, Value: fmt.Sprint(user.Flags), Inline: inline},
				{Name: `Public flags`, Value: fmt.Sprint(user.PublicFlags), Inline: inline},
				{Name: `Premium type`, Value: fmt.Sprint(user.PremiumType), Inline: inline},
				{Name: `Deaf`, Value: check(fmt.Sprint(member.Deaf)), Inline: inline},
				{Name: `GuildID`, Value: check(member.GuildID), Inline: inline},
				{Name: `JoinedAt`, Value: check(fmt.Sprint(member.JoinedAt)), Inline: inline},
				{Name: `Mute`, Value: check(fmt.Sprint(member.Mute)), Inline: inline},
				{Name: `Nick`, Value: check(member.Nick), Inline: inline},
				{Name: `Roles`, Value: check(fmt.Sprint(member.Roles)), Inline: inline},
			},
			Author: &dg.MessageEmbedAuthor{
				Name:    m.Author.Username,
				IconURL: m.Author.AvatarURL(""),
			},
		}

		if _, err := s.ChannelMessageSendEmbed(m.ChannelID, msg); err != nil {
			panic(err)
		}
	}
}

func check(s string) string {
	if len(s) > 0 {
		return s
	}
	return "None"
}
