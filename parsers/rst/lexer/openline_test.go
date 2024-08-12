package lexer

import (
	"testing"

	"github.com/avoceteditors/avocet/parsers/rst/token"
)

func Test_openline(t *testing.T) {
	name := "rst/lexer.Test_openline"
	data := []struct {
		s string
		t token.Type
		a Actus
		r int
	}{
		{"Text", token.ParaStart, ParaFindEnd, 4},
		{"\n", token.Empty, OpenLine, 1},
		{" ", token.Empty, OpenLine, 1},
		{" T", token.Indent, OpenLine, 2},
	}

	for pos, datum := range data {
		l := New("openline_test.go", datum.s)
		tok := l.Next()

		if tok.Type != datum.t {
			t.Fatalf("%s(%d): token returned wrong type: %q != %q", name, pos, tok.Type.String(), datum.t.String())
		}
		if l.actus != datum.a {
			t.Fatalf("%s(%d): actus in wrong state: %q != %q", name, pos, l.actus.String(), datum.a.String())
		}
		for i := 0; i < datum.r; i++ {
			tok = l.Next()
		}
		if tok.Type != token.EOF {
			t.Fatalf("%s(%d): token type not EOF: %q != %q", name, pos, tok.Type.String(), token.EOF.String())
		}
	}

}
