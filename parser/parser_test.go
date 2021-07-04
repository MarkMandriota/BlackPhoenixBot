// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const magic = `has "this line :3" &magic_parser`

func TestParser_Parse(t *testing.T) {
	dst := Routine{}
	NewFastParser(strings.NewReader(magic)).NextRoutine(&dst)

	t.Log(dst)
	if assert.Equal(t, Routine{
		Name: "magic_parser",
		Args: []Token{
			{DW, "has"},
			{DS, "this line :3"},
		},
	}, dst) {
		t.Log("Magic parser is not alone! >3")
	} else {
		t.Log("This line is just a dream of a magic parser, but dreams of a magic parser someday become real..")
	}
}

func BenchmarkParser_Parse(b *testing.B) {
	r := new(strings.Reader)
	p := NewFastParser(r)

	c := &Routine{Args: make([]Token, 0, 1<<0xB)}
	for i := 0; i < b.N; i++ {
		r.Reset(magic)
		_ = p.NextRoutine(c)
	}
}
