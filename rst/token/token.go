package token

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

	return tok
}
