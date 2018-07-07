package scanner

import (
	"fmt"
	"io"
	"os"
	"regexp"

	ts "text/scanner"

	"toolman.org/timespan"
)

type Scanner struct {
	mode     uint
	timespan *timespan.Timespan
	regex    *regexp.Regexp
	keyword  *KeyWord
	kwInfo   *keywordInfo
	doubles  map[rune]rune
	pairs    map[rune]*RunePair
	labels   map[rune]string
	text     string
	gs       *ts.Scanner
}

// Source is a source for Scanner text input. Any io.Reader having a
// 'Name() string' method can be used as a Source.
// (Not so) coincidentally, *os.File is a valid Source.
type Source interface {
	Name() string
	io.Reader
}

func New(src Source, options ...Option) *Scanner {
	gs := new(ts.Scanner).Init(src)

	gs.Mode = ts.GoTokens
	gs.Filename = src.Name()
	gs.Whitespace = ts.GoWhitespace | 1<<'\v' | 1<<'\f'

	s := &Scanner{
		gs:     gs,
		mode:   gs.Mode,
		labels: make(map[rune]string),
	}

	for _, opt := range options {
		opt(s)
	}

	return s
}

func (s *Scanner) error(msg string) {
	if s.gs.Error != nil {
		s.gs.Error(s.gs, msg)
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", s.Position(), msg)
}

func (s *Scanner) Scan() rune {
	tok := s.gs.Scan()
	s.text = s.gs.TokenText()
	s.timespan = nil
	s.regex = nil
	s.keyword = nil

	switch tok {
	case '#':
		if s.mode&ScanHashComments == 0 {
			return tok
		}
		return s.scanHashComment()

	case Ident:
		if ktok, ok := s.scanKeywords(tok); ok {
			return ktok
		}
		return tok

	case Int:
		if s.mode&ScanTimespans == 0 {
			return tok
		}
		return s.scanTimespan(tok)

	case '/':
		if s.mode&ScanRegexen == 0 {
			return tok
		}
		return s.scanRegex()

	default:
		if ptok, ok := s.scanPairs(tok); ok {
			return ptok
		}

		if dtok, ok := s.scanDoubles(tok); ok {
			return dtok
		}

		return tok
	}
}

func (s *Scanner) TokenString(tok rune) string {
	if ts, ok := s.labels[tok]; ok {
		return ts
	}

	return TokenString(tok)
}

var tokenStrings = make(map[rune]string)

func TokenString(tok rune) string {
	if s, ok := tokenStrings[tok]; ok {
		return s
	}
	return ts.TokenString(tok)
}

func (s *Scanner) Peek() rune   { return s.gs.Peek() }
func (s *Scanner) Text() string { return s.text }

// Returns the position of the most recently scanned token or, if that is
// invalid, the position of the character immediately following the most
// recently scanned token or character.
func (s *Scanner) Position() ts.Position {
	if s.gs.Position.IsValid() {
		return s.gs.Position
	}
	return s.gs.Pos()
}
