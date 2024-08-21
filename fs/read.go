package fs

import (
	"os"
)

// ReadFile takes a string filename and returns a rune channel.
// Internally, it starts a goroutine that passes the file
// content as runes to the channel.
// If the file cannot be read, it returns a closed channel.
func ReadFile(file string) chan rune {
	initLogger()
	con, err := os.ReadFile(f)
	if err != nil {
		lgr.Errorf("Read Error: %s\n%s", file, err)
		ch := make(chan rune, 1)
		close(ch)
		return ch
	}
	return runeHandler([]rune(string(con)))
}
