// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

package parser

import (
	"strings"
	"testing"
)

func TestIter_Next(t *testing.T) {
	iter := NewIter(NewParser(strings.NewReader(`:bar &foo :foobar &bar :foo &foobar`)))

	var p Pair
	for iter.Next(&p) {
		t.Log(p)
	}
}