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

// FastMagicParser - fast implementation of MagicParser.
type FastMagicParser struct {
	FI io.ByteScanner

	cc byte
	bf strings.Builder
}

// NewFastMagicParser - inits with fi and returns new FastMagicParser.
func NewFastMagicParser(fi io.ByteScanner) *FastMagicParser {
	return &FastMagicParser{FI: fi}
}

// NextRoutine - implements MagicParser.NextRoutine.
//go:nosplit
func (m *FastMagicParser) NextRoutine(r *Routine) bool {
	if r == nil {
		panic("nil pointer")
	}

	r.Name, r.Args = r.Name[:0], r.Args[:0]

	for m.nextB() != 0 {
		switch m.cc {
		case ' ', '\n', '\r', '\t', '\v':
		case '#':
			m.skipComment()
		case '&':
			r.Name = m.nextW()
			return true
		case ':':
			r.Args = append(r.Args, Token{
					T: ID,
					V: m.nextW(),
			})
		case '"':
			r.Args = append(r.Args, Token{
				T: DS,
				V: m.nextS(),
			})
		default:
			m.FI.UnreadByte()

			switch {
			case isDigit(m.cc):
				fallthrough
			case isLetter(m.cc):
				r.Args = append(r.Args, Token{
					T: DW,
					V: m.nextW(),
				})
			}
		}
	}

	return false
}

func (m *FastMagicParser) skipComment() {
	for m.nextB() != '#' {}
}

func (m *FastMagicParser) nextB() byte {
	m.cc, _ = m.FI.ReadByte()
	return m.cc
}

func (m *FastMagicParser) nextW() string {
	m.bf.Reset()

	for m.nextB() != 0 && (isLetter(m.cc) || isDigit(m.cc)) {
		m.bf.WriteByte(m.cc)
	}

	return m.bf.String()
}

func (m *FastMagicParser) nextS() string {
	m.bf.Reset()

	for m.nextB() != 0 && m.cc != '"' {
		if m.cc == '\\' {
			switch m.nextB() | 0x20 {
			case '"', '\\':
				m.bf.WriteByte(m.cc)
			default:
				m.FI.UnreadByte()
			}
			continue
		}
		m.bf.WriteByte(m.cc)
	}

	return m.bf.String()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' || c == '_'
}
