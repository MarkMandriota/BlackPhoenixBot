package token

type TokenType uint16
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = iota
	IDENT
	STR
	NUM

	EQU
)
