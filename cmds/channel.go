package cmds

import (
	dg "github.com/bwmarrin/discordgo"
)

type channel struct{}

func (*channel) Info(a ...string) string {
	if len(a) > 0 {
		switch a[0] {
		case "sub":
			return `Usage: <prefix>channel sub <channels name>
					Deletes channels`
		}
	}

	return `Usage: <prefix>channel <operation: add | sub>
			Manage channels`
}

func (*channel) Perm(a ...string) int64 {
	return dg.PermissionManageChannels
}

func (*channel) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	if len(a) > 1 {
		switch a[1] {
		case "sub":
			channels, _ := s.GuildChannels(m.GuildID)
			for _, name := range a[2:] {
				for _, channel := range channels {
					if channel.Name == name {
						s.ChannelDelete(channel.ID)
					}
				}
			}
		}
	}
}
