// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.


package parser

import (
	"testing"
	"unsafe"
)

func TestToken(t *testing.T) {
	t.Log(unsafe.Sizeof(Token{}))
}
