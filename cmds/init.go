package cmds

import (
	dg "github.com/bwmarrin/discordgo"
)

type Command interface {
	Info() string
	Exec(s *dg.Session, m *dg.MessageCreate)
}

var List map[string]Command

func init() {
	List = map[string]Command{
		"ping": new(ping),
	}
}
