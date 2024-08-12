package lexer

import "github.com/avoceteditors/avocet/parsers/rst/token"

func (l *Lexer) analyzeOpen(tok *token.Token) {
	switch l.get(l.pos) {
	case 0:
		tok.Type = token.EOF
	case '\n':
		tok.Type = token.Empty
		l.NL()
	case ' ', '\t':
		l.analyzeIndent(tok)
	case '*', '+', '-':
		l.analyzeUList(tok)
	default:
		tok.Type = token.ParaStart
		l.actus = ParaFindEnd
		l.pos--
	}
}

func (l *Lexer) analyzeIndent(tok *token.Token) {
	tok.Indent = l.getIndent(l.pos)
	l.pos += tok.Indent
	if l.get(l.pos) == '\n' || l.get(l.pos) == 0 {
		l.NL()
		tok.Type = token.Empty
		return
	}
	tok.Type = token.Indent
}

func (l *Lexer) getIndent(pos int) int {
	space := 0
	r := l.get(pos)
	for r == ' ' || r == '\t' {
		space += incSpace(r)
		pos++
		r = l.get(pos)
	}
	return space
}

func incSpace(r rune) int {
	if r == ' ' {
		return 1
	}
	return 3
}

// NL increments the line number and resets the column number
func (l *Lexer) NL() {
	l.line++
	l.col = 0
}
