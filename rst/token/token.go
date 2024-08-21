package token

import "fmt"

// Token struct represents a rune taken from a file or string.
// Once initialized it's used to perform lexical analysis and
// then passed thence to the parser.
type Token struct {
	Type Type `json:"type" yaml:"type" bson:"type"`

	File string `json:"file_name" yaml:"file" bson:"file"`
	Pos  int    `json:"file_pos" yaml:"pos" bson:"pos"`
	Line int    `json:"file_line" yaml:"line" bson:"line"`
	Col  int    `json:"file_col" yaml:"col" bson:"col"`

	Rune  rune   `json:"rune" yaml:"rune" bson:"rune"`
	Runes []rune `json:"runes" yaml:"runes" bson:"runes"`
	Size  int    `json:"size" yaml:"size" bson:"size"`
}

// New takes a filename, line, column, position, and rune
// and initializes a Token pointer, which it returns to the
// caller.
func New(fname string, line, col, pos int, r rune) *Token {
	tok := &Token{
		Type:  Text,
		File:  fname,
		Pos:   pos,
		Line:  line,
		Col:   col,
		Rune:  r,
		Runes: []rune{},
	}
	if r != 0 {
		tok.Append(r)
	}

	return tok
}

// String method returns a string representation of the Token.
func (tok *Token) String() string {
	return fmt.Sprintf(
		"Token<%s>{file=%q, line=%d, col=%d, text=%q}",
		tok.Type.String(),
		tok.File, tok.Line, tok.Col, string(tok.Runes))
}

// Append takes a rune and appends it to the interal runes
// stored in the Token.  This is used to build out runes
// by the lexer when grouping a related set.
func (tok *Token) Append(r rune) {
	tok.Runes = append(tok.Runes, r)
	tok.Size++
}
