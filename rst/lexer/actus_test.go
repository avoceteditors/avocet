package lexer

import "testing"

func Test_actus(t *testing.T) {
	name := "rst/lexer.Test_actus"
	data := []struct {
		t Actus
		s string
	}{
		{OpenLine, openLine},
		{PassAct, passactLine},
		{ParaFindEnd, paraFindEnd},
	}

	for pos, datum := range data {
		s := datum.t.String()

		if s != datum.s {
			t.Fatalf("%s(%d): actus state returned wrong string: %q != %q", name, pos, s, datum.s)
		}
	}
}
