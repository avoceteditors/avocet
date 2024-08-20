package fs

import (
	"testing"
	"time"
)

func Test_stringReader(t *testing.T) {
	name := "fs.Test_stringReader"
	s := "Test Read String"
	rs := []rune(s)
	ch := ReadString(s)
	pos := 0

	for pos < len(s) {
		select {
		case <-time.After(1 * time.Millisecond):
			t.Fatalf("%s: Timeout", name)
		case r := <-ch:
			if r != rs[pos] {
				t.Errorf("%s: Expected %c, got %c", name, s[pos], r)
			}
			pos++
		}
	}

}
