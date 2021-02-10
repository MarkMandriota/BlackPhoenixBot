package cmds

import (
	"fmt"
	"strconv"

	dg "github.com/bwmarrin/discordgo"
)

type clear struct{}

func (*clear) Info(...string) string {
	return `Usage: <prefix>clear <?number>
			Clear certain number of message`
}

func (*clear) Exec(s *dg.Session, m *dg.MessageCreate, a ...string) {
	perm, err := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x", perm)
	if perm&0x00002000 != 0x00002000 && perm&0x00000008 != 0x00000008 {
		return
	}

	num := 100
	if len(a) > 1 {
		num, err = strconv.Atoi(a[1])
		if err != nil {
			panic(&dg.MessageEmbed{
				Color:       0xFF0000,
				Title:       "ERROR",
				Description: "Argument has not integer type",
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
