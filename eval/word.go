package eval

import (
	"regexp"
	"sync"
)

var rep = regexp.MustCompile(`\b(\w+)\s+\1\b`)

func (e *Evaluator) wordWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		if rep.MatchString(l.Text) {
			l.Error("Dup")
		}
	}
}
