package token

import "fmt"

// Token represents a lexical token
type Token struct {
	Type Type   `json:"type" yaml:"type"`
	Text string `json:"text" yaml:"text"`
	Raw  []rune `json:"raw" yaml:"raw"`
	Size int    `json:"size" yaml:"size"`

	File   string `json:"file" yaml:"file"`
	Line   int    `json:"line" yaml:"line"`
	Col    int    `json:"col" yaml:"col"`
	Indent int    `json:"indent" yaml:"indent"`
}

// New creates a new token
func New(text, fname string, line, col int) *Token {
	tok := &Token{
		Type: Unset,
		Text: text,
		Raw:  []rune(text),

		File: fname,
		Line: line,
		Col:  col,
	}
	tok.Size = len(tok.Raw)

	return tok
}

func (tok *Token) String() string {
	return fmt.Sprintf("Token<%s>", tok.Type.String())
}

func (tok *Token) Get(i int) rune {
	if i < tok.Size {
		return tok.Raw[i]
	}
	return 0
}

// Append adds a new rune to the token.
func (tok *Token) Append(r rune) {
	tok.Raw = append(tok.Raw, r)
	tok.Size++
}

// Render converts the raw rune slice to a string and sets it on Token.Text.
func (tok *Token) Render() {
	tok.Text = string(tok.Raw)
}
