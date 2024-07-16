package eval

import (
	"fmt"
	"sync"
	"unicode"
)

func (e *Evaluator) wordWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		rs := []rune{}

		ls := []string{}
		for _, c := range l.Text {
			c = unicode.ToLower(c)
			if unicode.IsLetter(c) {
				rs = append(rs, c)
			} else if len(rs) > 0 {
				ls = append(ls, string(rs))
				rs = []rune{}
			}
		}
		lastWord := ""
		for _, w := range ls {
			if w == lastWord {
				l.Error(fmt.Sprintf("Dup(%q)", lastWord))
				e.hasErrors = 1
			}
			lastWord = w
		}

	}
}
