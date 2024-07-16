package eval

import (
	"sync"
	"unicode"
)

func (e *Evaluator) wordWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		lastWord := ""
		for _, w := range l.Words {
			nw := []rune{}
			for _, r := range w {
				if unicode.IsLetter(r) || unicode.IsNumber(r) {
					nw = append(nw, r)
				}
			}
			ns := string(nw)
			if lastWord == ns {
				l.Error("Dup")
				e.hasErrors = 1
			}
			lastWord = ns

		}
	}
}
