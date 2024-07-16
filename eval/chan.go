package eval

func (e *Evaluator) Sends(lcs []chan *Line) {
	for _, p := range e.Paths {
		e.SendLines(p, lcs)
	}
}

func (e *Evaluator) Close(lcs []chan *Line) {
	for _, c := range lcs {
		close(c)
	}
}
