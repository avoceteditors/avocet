package lexer

import (
	"github.com/avoceteditors/avocet/rst/token"
)

func (l *Lexer) analyze(tok *token.Token) {
	switch l.actus {
	case OpenLine:
		l.analyzeOpen(tok)
	}
}
