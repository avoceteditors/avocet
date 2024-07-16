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
		last := rune(0)
		for _, c := range l.Text {
			if c == ' ' && len(rs) > 0 {
				ls = append(ls, string(rs))
				rs = []rune{}
			} else {
				rs = append(rs, unicode.ToLower(c))
			}
			last = c
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
