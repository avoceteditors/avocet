package token

import "testing"

func Test_analysis(t *testing.T) {
	name := "rst/token.Test_analysis"
	data := []struct {
		r rune
		t Type
		l int
	}{
		{'*', Star, 1},
		{'`', Tick, 1},
		{':', Colon, 1},
		{' ', Space, 1},
		{'\t', Space, 3},
		{'\n', EOL, 1},
		{'.', Dot, 1},
		{'[', OBrack, 1},
		{']', CBrack, 1},
		{'<', OABrack, 1},
		{'>', CABrack, 1},
		{'~', Tilda, 1},
		{'_', Score, 1},
		{'+', Plus, 1},
		{'0', Num, 1},
	}

	f := "analyze_test.go"
	for pos, datum := range data {
		tok := New(f, 1, pos, pos, datum.r)
		tok.Analyze()

		if tok.Rune != datum.r {
			t.Errorf("%s(%d): token set wrong rune: %q != %q", name, pos, string(tok.Rune), string(datum.r))
		}

		if tok.Type != datum.t {
			t.Errorf("%s(%d): token set wrong type: %q != %q", name, pos, tok.Type.String(), datum.t.String())
		}
		if tok.Size != datum.l {
			t.Errorf("%s(%d): token set wrong length: %d != %d", name, pos, tok.Size, datum.l)
		}

	}

}
