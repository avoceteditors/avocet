package eval

import (
	"regexp"
	"sync"
)

var dspace = regexp.MustCompile(`\.  `)

func (e *Evaluator) spaceWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		if dspace.MatchString(l.Text) {
			l.Error("Space")
			e.hasErrors = 1
		}
	}
}
