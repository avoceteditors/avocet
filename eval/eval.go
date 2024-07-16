package eval

import (
	"runtime"
	"sync"

	"github.com/avoceteditors/avocet/ast"
)

type Evaluator struct {
	Paths     []*ast.Path
	Size      int
	hasErrors int
}

func New(ps []*ast.Path) *Evaluator {
	e := &Evaluator{
		Paths: ps,
		Size:  len(ps),
	}
	return e
}

func (e *Evaluator) Run() int {

	spaceIn := make(chan *Line, e.Size*4)
	wordIn := make(chan *Line, e.Size*4)
	var wg sync.WaitGroup

	for w := 0; w < runtime.NumCPU()*4; w++ {
		go e.spaceWorker(spaceIn, wg)
		go e.wordWorker(wordIn, wg)
	}

	// Send Data to Channels
	lcs := []chan *Line{spaceIn, wordIn}
	e.Sends(lcs)
	e.Close(lcs)

	// Wait til Done
	wg.Wait()
	return e.hasErrors
}
