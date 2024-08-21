package token

import "testing"

func Test_types(t *testing.T) {
	name := "rst/token.Test_types"
	data := []struct {
		t Type
		s string
	}{
		{Text, textType},
		{EOF, eofType},
		{EOL, eolType},
		{Star, starType},
		{Tick, tickType},
		{Colon, colonType},
		{Space, spaceType},
		{Dot, dotType},
		{OBrack, obrackType},
		{CBrack, cbrackType},
		{OABrack, oabrackType},
		{CABrack, cabrackType},
		{Tilda, tildaType},
		{Score, scoreType},
		{Plus, plusType},
	}

	for pos, datum := range data {
		s := datum.t.String()

		if s != datum.s {
			t.Errorf(
				"%s(%d): type returned wrong string: %q != %q",
				name, pos, s, datum.s)
		}
	}
}
