package token

import "testing"

func Test_types(t *testing.T) {
	name := "rst/token.Test_types"
	data := []struct {
		t Type
		s string
	}{
		{Unset, unsetType},
		{EOL, eolType},
		{EOF, eofType},
		{ParaStart, paraStartType},
		{ParaEnd, paraEndType},
		{Empty, emptyType},
		{UList, ulistType},
		{OList, olistType},
		{CList, clistType},
		{Indent, indentType},
	}

	for pos, datum := range data {
		s := datum.t.String()

		if s != datum.s {
			t.Fatalf("%s(%d): token type returned wrong string: %q != %q", name, pos, s, datum.s)
		}
	}
}
