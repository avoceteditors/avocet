package lexer

import "github.com/avoceteditors/avocet/rst/token"

type Lexer struct {
	actus Actus

	text string
	raw  []rune
	size int

	file   string
	line   int
	col    int
	pos    int
	indent int
}

// New creates a new lexer
func New(fname, text string) *Lexer {
	l := &Lexer{
		file: fname,
		text: text,
		raw:  []rune(text),
		line: 1,
		pos:  -1,
	}
	l.size = len(l.raw)
	return l
}

func (l *Lexer) get(i int) rune {
	if i < l.size {
		return l.raw[i]
	}
	return 0
}

// Next returns the next token
func (l *Lexer) Next() *token.Token {
	l.pos++
	l.col++
	tok := token.New("", l.file, l.line, l.col)
	if l.pos >= l.size {
		tok.Type = token.EOF
		return tok
	}
	l.analyze(tok)

	return tok
}
