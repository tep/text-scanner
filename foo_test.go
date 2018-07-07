package scanner

import (
	"fmt"
	"os"
	"testing"
)

const (
	DOIT = iota + 2345
	BRACES
	ASSIGN
	LDANGLE
	RDANGLE
)

func TestFoo(t *testing.T) {
	f, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}

	modes := Modes(+ScanHashComments, -SkipComments, +ScanRegexen)
	kwrds := Keywords(&KeyWord{"doit", DOIT})
	pairs := Pairs(&RunePair{'{', '}', BRACES}, &RunePair{':', '=', ASSIGN})
	dubls := Doubles(&Double{'<', LDANGLE}, &Double{'>', RDANGLE})

	s := New(f, modes, kwrds, pairs, dubls, Error(func(s *Scanner, msg string) { err = fmt.Errorf("%s: %v", msg, s.Position()) }))

	var tok rune

	for tok != EOF {
		tok = s.Scan()

		if err != nil {
			t.Logf("ERROR: %v\n", err)
			err = nil
			continue
		}

		t.Logf("%s: %q\n", s.TokenString(tok), s.Text())
	}
}
