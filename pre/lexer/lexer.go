package lexer

type Lexer struct {
	in     string
	ch     rune
	pp, np int
}

func NewLexer() *Lexer {
	l := &Lexer{}
	return l
}

func (l *Lexer) Lex() {

}
