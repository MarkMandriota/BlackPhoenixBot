// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"io"
	"strings"
)

// Routine - result struct of magic parsing.
type Routine struct {
	Name string
	Args []Token
}

// String - implements fmt.Stringer.
func (c *Routine) String() (s string) {
	s += fmt.Sprintf("#%s\n", c.Name)

	for i, arg := range c.Args {
		s += fmt.Sprintf("\t%d: T=%d V=%s\n", i, arg.T, arg.V)
	}
	return
}

// MagicParser - iterator struct of magic source text.
type MagicParser interface {
	// NextRoutine - reads next routine into struct. r must be not nil.
	NextRoutine(r *Routine) bool
}

// Parser - fast implementation of MagicParser.
type Parser struct {
	FI io.ByteScanner

	cc byte
	bf strings.Builder
}

// NewParser - inits with fi and returns new FastMagicParser.
func NewParser(fi io.ByteScanner) *Parser {
	return &Parser{FI: fi}
}

// NextRoutine - implements MagicParser.NextRoutine.
//go:nosplit
func (p *Parser) NextRoutine(r *Routine) bool {
	if r == nil {
		panic("nil pointer")
	}

	r.Name, r.Args = r.Name[:0], r.Args[:0]

	for p.nextB() != 0 {
		switch p.cc {
		case ' ', '\n', '\r', '\t', '\v':
		case '#':
			p.skipComment()
		case '&':
			r.Name = p.nextW()
			return true
		case ':':
			r.Args = append(r.Args, Token{
					T: ID,
					V: p.nextW(),
			})
		case '"':
			r.Args = append(r.Args, Token{
				T: DS,
				V: p.nextS(),
			})
		default:
			p.FI.UnreadByte()

			switch {
			case isDigit(p.cc):
				fallthrough
			case isLetter(p.cc):
				r.Args = append(r.Args, Token{
					T: DW,
					V: p.nextW(),
				})
			}
		}
	}

	return false
}

func (p *Parser) skipComment() {
	for p.nextB() != '#' {}
}

func (p *Parser) nextB() byte {
	p.cc, _ = p.FI.ReadByte()
	return p.cc
}

func (p *Parser) nextW() string {
	p.bf.Reset()

	for p.nextB() != 0 && (isLetter(p.cc) || isDigit(p.cc)) {
		p.bf.WriteByte(p.cc)
	}

	return p.bf.String()
}

func (p *Parser) nextS() string {
	p.bf.Reset()

	for p.nextB() != 0 && p.cc != '"' {
		if p.cc == '\\' {
			switch p.nextB() | 0x20 {
			case '"', '\\':
				p.bf.WriteByte(p.cc)
			default:
				p.FI.UnreadByte()
			}
			continue
		}
		p.bf.WriteByte(p.cc)
	}

	return p.bf.String()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' || c == '_'
}
