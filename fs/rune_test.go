package fs

import (
	"testing"
	"time"
)

func Test_runeHandler(t *testing.T) {
	name := "fs.Test_runeHandler"
	rs := []rune("Test")
	ch := runeHandler(rs)
	pos := 0

	for pos < len(rs) {
		select {
		case <-time.After(1 * time.Millisecond):
			t.Fatalf("%s: Timeout", name)
		case r := <-ch:
			if r != rs[pos] {
				t.Errorf("%s: Expected %c, got %c", name, rs[pos], r)
			}
			pos++
		}
	}

}
