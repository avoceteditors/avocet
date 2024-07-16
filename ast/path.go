package ast

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type Path struct {
	Rel   string    `json:"path_relative" yaml:"path_relative" bson:"path_relative"`
	Abs   string    `json:"path_absolute" yaml:"path_absolute" bson:"path_absolute"`
	Name  string    `json:"name" yaml:"name" bson:"name"`
	Fname string    `json:"filename" yaml:"filename" bson:"filename"`
	Dir   string    `json:"parent" yaml:"parent" bson:"parent"`
	Ext   string    `json:"extension" yaml:"extension" bson:"extension"`
	Mtime time.Time `json:"mtime" yaml:"mtime" bson:"mtime"`
	IsDir bool      `json:"-" yaml:"-" bson:"-"`
}

func NewPath(arg string) *Path {
	// Initialize Path
	p := &Path{
		Abs:   resolvePath(arg),
		Fname: filepath.Base(arg),
		Ext:   filepath.Ext(arg),
	}

	// Set Name
	p.Name = p.Fname
	if p.Ext != "" {
		p.Name = strings.TrimSuffix(p.Fname, p.Ext)
	}
	p.Dir = filepath.Dir(p.Abs)

	// Set Relative Path
	var err error
	p.Rel, err = filepath.Rel(viper.GetString("paths.wd"), p.Abs)
	if err != nil {
		log.Warnf("Error setting relative paths, defaulting to absolute: %s", err)
		p.Rel = p.Abs
	}

	// Stat
	fi, err := os.Stat(p.Abs)
	if err != nil {
		log.Warnf("Unable to stat file: %s", err)
		return p
	}
	p.Mtime = fi.ModTime()
	p.IsDir = fi.IsDir()

	return p
}

func (p *Path) ReadDirs() []*Path {
	ps := []*Path{}
	if !p.IsDir {
		return append(ps, p)
	}
	fs, err := os.ReadDir(p.Abs)
	if err != nil {
		log.Error("Unable to read directory: %s\n%s", p.Rel, err)
		return ps
	}
	for _, fi := range fs {
		f := filepath.Join(p.Abs, fi.Name())
		np := NewPath(f)
		ps = append(ps, np.ReadDirs()...)
	}

	return ps
}

func resolvePath(arg string) string {
	rs := []rune(arg)
	size := len(rs)
	hm := viper.GetString("paths.home")
	wk := viper.GetString("paths.wd")
	switch size {
	case 0:
		return wk
	case 1:
		switch rs[0] {
		case '~':
			return hm
		case '.':
			return wk

		}
	case 2:
		if rs[1] == '/' {
			switch rs[0] {
			case '~':
				return hm
			case '.':
				return wk
			}
		}
	default:
		switch rs[0] {
		case '~':
			if rs[1] == '/' {
				filepath.Join(hm, string(rs[2:]))
			}
		case '/':
			return arg
		case '.':
			if rs[1] == '.' && rs[2] == '/' {
				return filepath.Join(filepath.Dir(wk), string(rs[3:]))
			} else if rs[1] == '/' {
				return filepath.Join(wk, string(rs[2:]))
			}
		default:
			return filepath.Join(wk, arg)
		}
	}
	return arg
}

func (p *Path) GetLines() []string {
	ls := []string{}
	con, err := os.ReadFile(p.Abs)
	if err != nil {
		log.Warnf("Unable to read file %s:\n%s", p.Rel, err)
		return ls
	}
	return strings.Split(string(con), "\n")

}

var wspaces = regexp.MustCompile(" +")

func (p *Path) GetWords(ls []string) []string {
	ws := []string{}
	con, err := os.ReadFile(p.Abs)
	if err != nil {
		log.Warnf("Unable to read file %s:\n%s", p.Rel, err)
		return ws
	}
	return wspaces.Split(string(con), -1)
}
