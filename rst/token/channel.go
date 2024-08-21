package token

import "github.com/spf13/viper"

// Tokenize takes a file string and a channel that receives runes
// and returns a channel that receives Token pointers.
func Tokenize(file string, in <-chan rune) chan *Token {
	out := make(chan *Token, viper.GetInt("avocet.buffer.token")+1)
	go func() {
		line := 1
		col := 1
		pos := 0
		for r := range in {
			tok := New(file, line, col, pos, r)
			tok.Analyze()

			// Increment Line/Column/Position
			if tok.Type == EOL {
				line++
				col = 0
			}
			col++
			pos++

			// Send Token
			out <- tok
		}

		// Send EOF Token
		tok := New(file, line, col, pos, 0)
		tok.Type = EOF
		out <- tok
		close(out)
	}()

	return out
}
