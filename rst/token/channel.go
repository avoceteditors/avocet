package token

import "github.com/spf13/viper"

func Tokenize(file string, in chan rune) chan *Token {
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
