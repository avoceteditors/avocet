package fs

import (
	"os"
)

func ReadFile(f string) chan rune {
	con, err := os.ReadFile(f)
	if err != nil {
		lgr.Errorf("Read Error: %s\n%s", f, err)
		ch := make(chan rune, 1)
		close(ch)
		return ch
	}
	return runeHandler([]rune(string(con)))
}
