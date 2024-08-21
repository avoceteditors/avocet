package token

import "unicode"

// Analyze method classifies the initial token rune to set
// a type.  This prepares the rune for lexical analysis.
func (tok *Token) Analyze() {
	switch tok.Rune {
	case '*':
		tok.Type = Star
	case '`':
		tok.Type = Tick
	case ':':
		tok.Type = Colon
	case ' ':
		tok.Type = Space
	case '\t':
		tok.Type = Space
		tok.Runes = []rune("   ")
		tok.Size = 3
	case '\n':
		tok.Type = EOL
	case '.':
		tok.Type = Dot
	case '[':
		tok.Type = OBrack
	case ']':
		tok.Type = CBrack
	case '<':
		tok.Type = OABrack
	case '>':
		tok.Type = CABrack
	case '~':
		tok.Type = Tilda
	case '_':
		tok.Type = Score
	case '+':
		tok.Type = Plus
	default:
		if unicode.IsNumber(tok.Rune) {
			tok.Type = Num
		}
	}
}
