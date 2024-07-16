package eval

import (
	"strings"
	"sync"
	"unicode"
)

func (e *Evaluator) wordWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		rs := []rune{}
		var lastChar rune

		for _, c := range l.Text {
			if unicode.IsLetter(c) {
				rs = append(rs, c)
			} else if c == ' ' && lastChar != ' ' {
				rs = append(rs, ' ')
			}
			lastChar = c
		}
		ls := strings.Split(string(rs), " ")
		lastWord := ""
		for _, w := range ls {
			if w == lastWord {
				l.Error("Dup")
				e.hasErrors = 1
			}
			lastWord = w
		}

	}
}
