package eval

import (
	"fmt"
	"regexp"

	"github.com/avoceteditors/avocet/ast"
)

type Line struct {
	Rel   string
	Line  int
	Text  string
	Words []string
}

var sp = regexp.MustCompile(" +")

func NewLine(rel string, line int, text string) *Line {
	return &Line{
		Rel:   rel,
		Line:  line,
		Text:  text,
		Words: sp.Split(text, -1),
	}
}

func (l *Line) Error(e string) {
	fmt.Printf("%s: %s:%d: %s\n", e, l.Rel, l.Line, l.Text)
}

func (e *Evaluator) SendLines(p *ast.Path, lsc []chan *Line) {
	ls := p.GetLines()
	for n, s := range ls {
		l := NewLine(p.Rel, n, s)
		for _, c := range lsc {
			c <- l
		}
	}
}
