package lexer

import "github.com/avoceteditors/avocet/rst/token"

func (l *Lexer) analyzeUList(tok *token.Token) {
	prime := l.get(l.pos)
	secunde := l.get(l.pos + 1)
	terte := l.get(l.pos + 2)
	if prime == secunde && prime == terte {
		// Heading
		return
	}
	if secunde != ' ' {
		tok.Type = token.ParaStart
		l.pos--
		return
	}
	l.pos++
	l.getIndent()
	tok.Type = token.UList
}
