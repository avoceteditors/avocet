package token

import (
	"testing"

	"github.com/avoceteditors/avocet/fs"
)

func Test_channel(t *testing.T) {
	name := "rst/token.Test_channel"

	text := "hhh1`.:"
	ts := []Type{
		Text, Text, Text, Num, Tick, Dot, Colon,
	}
	size := len(ts)
	in := fs.ReadString(text)
	f := "channel_test.go"
	out := Tokenize(f, in)

	pos := 0
	for pos < size {
		tok := <-out
		if tok.Type != ts[pos] {
			t.Errorf("%s(%d): token set wrong type: %q != %q", name, pos, tok.Type.String(), ts[pos].String())
		}
		pos++
	}
}
