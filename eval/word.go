package eval

import "sync"

func (e *Evaluator) wordWorker(in <-chan *Line, wg sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for l := range in {
		last := ""
		for _, w := range l.Words {
			if w != "" {
				if last == w {
					l.Error("Duplicate")
				}
				last = w
			}
		}
	}
}
