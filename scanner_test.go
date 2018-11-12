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

	kwrds := KeywordMap{"doit": DOIT}

	pairs := RunePairs{
		{'{', '}', BRACES},
		{':', '=', ASSIGN},
	}

	dubls := Doubles{
		{'<', LDANGLE},
		{'>', RDANGLE},
	}

	var s *Scanner
	ef := ErrFunc(func(msg string) { err = fmt.Errorf("%s: %v", msg, s.Position()) })
	s = New(f, ef, kwrds, pairs, dubls, +ScanHashComments, +ScanTimespans, +ScanStdSizes, -SkipComments, +ScanRegexen)

	var tok rune

	for tok != EOF {
		tok = s.Scan()

		if err != nil {
			t.Logf("ERROR: %v\n", err)
			err = nil
			continue
		}

		var ts string
		switch tok {
		case KeyWord:
			ts = "Keyword"
		case BRACES:
			ts = "BRACES"
		case ASSIGN:
			ts = "ASSIGN"
		case LDANGLE:
			ts = "LDANGLE"
		case RDANGLE:
			ts = "RDANGLE"
		default:
			ts = s.TokenString(tok)
		}

		t.Logf("%s: %q\n", ts, s.Text())
	}
}
