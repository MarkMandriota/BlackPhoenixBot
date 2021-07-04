// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by a Apache 2.0-style
// license that can be found in the LICENSE file.

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
)