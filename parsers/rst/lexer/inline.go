package lexer

import (
	"unicode"

	"github.com/avoceteditors/avocet/parsers/rst/token"
)

func (l *Lexer) analyzeInline(tok *token.Token) {
	r := l.get(l.pos)
	if unicode.IsSpace(r) {
		tok.Type = token.Space
		l.skipNextSpace()
		return
	}
	switch r {
	default:
		tok.Raw = append(tok.Raw, r)
		for !unicode.IsSpace(l.get(l.pos + 1)) {
			l.pos++
			l.col++
			tok.Raw = append(tok.Raw, l.get(l.pos))
		}
		tok.Text = string(tok.Raw)
	}

}

func (l *Lexer) skipNextSpace() {
	for unicode.IsSpace(l.get(l.pos)) || l.get(l.pos) == '\n' {
		l.pos++
		l.col++
	}
}
