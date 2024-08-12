package avocet

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/avoceteditors/avocet/parsers/rst/lexer"
	"github.com/avoceteditors/avocet/parsers/rst/token"
	"github.com/spf13/viper"
)

func RunLexer(args []string) {
	if len(args) > 0 {
		switch viper.GetString("avocet.default_lexer") {
		case "rst":
			lexRst(strings.Join(args, " "))
		default:
			lexRst(strings.Join(args, " "))
		}
		return
	}
}

func lexRst(text string) {
	toks := []*token.Token{}
	l := lexer.New("STDIN", text)
	tok := l.Next()
	for tok.Type != token.EOF {
		toks = append(toks, tok)
		tok = l.Next()
	}
	toks = append(toks, tok)
	if len(toks) == 0 {
		fmt.Println("Error")
	}
	printToks(toks)
}

func printToks(toks []*token.Token) {
	con, err := json.Marshal(toks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(con))
}
