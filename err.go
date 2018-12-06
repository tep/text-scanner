package scanner

import (
	"fmt"
	"os"

	ts "text/scanner"
)

func (s *Scanner) error(msg string) {
	if s.gs.Error != nil {
		s.gs.Error(s.gs, msg)
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", s.Position(), msg)
}

type ErrFunc func(string)

func (ef ErrFunc) setOpt(s *Scanner) {
	s.gs.Error = func(_ *ts.Scanner, msg string) { ef(msg) }
}
