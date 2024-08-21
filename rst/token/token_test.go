package token

import "testing"

func Test_token_init(t *testing.T) {
	f := "token_test.go"
	l := 1
	col := 2
	p := 4
	r := 'a'

	tok := New(f, l, col, p, r)

	if tok.Type != Text {
		t.Errorf("Token has wrong type: %q != %q", textType, tok.Type.String())
	}
	if tok.Line != l {
		t.Errorf("Token has wrong line: %d != %d", l, tok.Line)
	}
	if tok.Col != col {
		t.Errorf("Token has wrong column: %d != %d", col, tok.Col)
	}
	if tok.Pos != p {
		t.Errorf("Token has wrong position: %d != %d", p, tok.Pos)
	}
	if tok.Rune != r {
		t.Errorf("Token has wrong rune: %q != %q", string(r), string(tok.Rune))
	}
	if len(tok.Runes) != 1 {
		t.Errorf("Token has wrong number of runes: %d != %d", 1, len(tok.Runes))
	}
}
