package fs

import (
	"os"
	"testing"
	"time"
)

func Test_fileRead(t *testing.T) {
	name := "fs.Test_fileRead"
	s := "Test Read File"
	f := "read_test.go"
	ch := ReadFile(f)
	con, err := os.ReadFile(f)
	rs := []rune(string(con))
	if err != nil {
		t.Fatalf("%s Unable to read file: %s\n%s", name, f, err)
	}

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
