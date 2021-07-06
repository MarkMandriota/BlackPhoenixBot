// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

package parser

// Pair - pair of 2 routines
type Pair [2]Routine

// MagicIter - iterator struct of Pair.
type MagicIter interface {
	// Next - returns next Pair.
	Next(p *Pair) bool
}

// Iter - implements MagicIter.
type Iter struct {
	p MagicParser
}

// NewIter - inits with p and returns new Iter.
func NewIter(p MagicParser) *Iter {
	return &Iter{p: p}
}

// Next - implements MagicIterator.Next.
func (i *Iter) Next(p *Pair) bool {
	p[0] = p[1]
	return i.p.NextRoutine(&p[1])
}