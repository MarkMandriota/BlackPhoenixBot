package kernel

import (
	dg "github.com/bwmarrin/discordgo"
)

type Bot struct {
	S      dg.Session
	Prefix string

	errs chan error
}
