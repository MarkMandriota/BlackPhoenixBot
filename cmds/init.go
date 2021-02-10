package cmds

import (
	dg "github.com/bwmarrin/discordgo"
)

type Command interface {
	Info(a ...string) string
	Exec(s *dg.Session, m *dg.MessageCreate, a ...string)
}

var Map map[string]Command

func init() {
	Map = map[string]Command{
		"help":  &help{},
		"ping":  &ping{},
		"send":  &send{},
		"info":  &info{},
		"clear": &clear{},
	}
}
