package cmds

import dg "github.com/bwmarrin/discordgo"

type rob struct{}

func (*rob) Info(...string) string {
	return `Usage: <prefix>rob <user> <role>
			steals roles`
}

func (*rob) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	if len(a) > 2 {
		if err := s.GuildMemberRoleRemove(m.GuildID, a[1], a[2]); err != nil {
			panic(&dg.MessageEmbed{
				Color:       0xFF0000,
				Title:       "ERROR",
				Description: "The bot does not have enough rights",
			})
		}
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, a[2])
	}
}
