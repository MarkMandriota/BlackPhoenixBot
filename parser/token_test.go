package parser

import (
	"testing"
	"unsafe"
)

func TestToken(t *testing.T) {
	t.Log(unsafe.Sizeof(Token{}))
}
