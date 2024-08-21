package token

import "fmt"

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

func (tok *Token) String() string {
	return fmt.Sprintf(
		"Token<%s>{file=%q, line=%d, col=%d, text=%q}",
		tok.Type.String(),
		tok.File, tok.Line, tok.Col, string(tok.Runes))
}

func (tok *Token) Append(r rune) {
	tok.Runes = append(tok.Runes, r)
	tok.Size++
}
