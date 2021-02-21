package cmds

import dg "github.com/bwmarrin/discordgo"

type rob struct{}

func (*rob) Info(...string) string {
	return `Usage: <prefix>rob <user> <role>
			steals roles`
}

func (*rob) Perm(a ...string) int64 {
	return dg.PermissionManageRoles
}

func (*rob) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	if len(a) > 2 {
		s.GuildMemberRoleRemove(m.GuildID, a[1], a[2])
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, a[2])
	}
}
