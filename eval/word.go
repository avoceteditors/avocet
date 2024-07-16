package eval

import (
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
			if unicode.IsLetter(c) || unicode.IsNumber(c) {
				rs = append(rs, c)
			} else if len(rs) > 0 {
				ls = append(ls, string(rs))
				rs = []rune{}
			}
		}
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
