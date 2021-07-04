// Copyright 2021 Mark Mandriota. All right reserved.

package parser

// Type of Token type
type Type uint

// Token - result struct of magic lexing.
type Token struct {
	T Type // Token type
	V string // Token value
}

// Token types
const (
	ID Type = iota // Identifier
	DW // Data Word
	DS // Data String
	DI // Data Integer
	DH // Data Hex
	DF // Data Float
)