package lexer

import "github.com/avoceteditors/avocet/parsers/rst/token"

func (l *Lexer) analyzePara(tok *token.Token) {
	switch l.get(l.pos) {
	case 0:
		tok.Type = token.ParaEnd
	case '\n':
		indent := l.getIndent(l.pos + 1)
		if indent == l.indent {
			l.analyzeInline(tok)
			return
		}
		tok.Type = token.ParaEnd
		l.actus = OpenLine
	default:
		l.analyzeInline(tok)
	}
}
