package cmds

import (
	dg "github.com/bwmarrin/discordgo"
)

type role struct{}

func (*role) Info(a ...string) string {
	if len(a) > 0 {
		switch a[0] {
		case "sub":
			return `Usage: <prefix>role sub <roles name>
					Deletes roles`
		}
	}

	return `Usage: <prefix>channel <operation: add | sub>
			Manage channels`
}

func (*role) Perm(a ...string) int64 {
	return dg.PermissionManageRoles
}

func (*role) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	if len(a) > 1 {
		switch a[1] {
		case "sub":
			roles, _ := s.GuildRoles(m.GuildID)
			for _, name := range a[2:] {
				for _, role := range roles {
					if role.Name == name {
						s.ChannelDelete(role.ID)
					}
				}
			}
		}
	}
}
