package avocet

import (
	"os"

	"github.com/avoceteditors/avocet/ast"
	"github.com/avoceteditors/avocet/eval"
	"github.com/charmbracelet/log"
)

func DocEval(args []string) {
	log.Info("Called the Documentation Evaluation Operation")
	ps := []*ast.Path{}
	for _, arg := range args {
		p := ast.NewPath(arg)
		ps = append(ps, p.ReadDirs()...)
	}
	log.Debugf("Found %d paths", len(ps))
	e := eval.New(ps)
	os.Exit(e.Run())

}
