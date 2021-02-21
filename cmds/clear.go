package cmds

import (
	"strconv"

	dg "github.com/bwmarrin/discordgo"
)

type clear struct{}

func (*clear) Info(...string) string {
	return `Usage: <prefix>clear <?number>
			Clear certain number of message`
}

func (*clear) Perm(a ...string) int64 {
	return dg.PermissionManageMessages
}

func (*clear) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	num := 100
	if err := error(nil); len(a) > 1 {
		num, err = strconv.Atoi(a[1])
		if err != nil {
			panic(&dg.MessageEmbed{
				Color:       0xFF0000,
				Title:       `ERROR`,
				Description: `Argument has not integer type`,
			})
		}
	}

	msg, err := s.ChannelMessages(m.ChannelID, num, "", "", "")
	if err != nil {
		panic(err)
	}

	id := make([]string, len(msg))
	for i, el := range msg {
		id[i] = el.ID
	}

	s.ChannelMessagesBulkDelete(m.ChannelID, id)
}
