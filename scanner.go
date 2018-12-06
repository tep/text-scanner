package scanner

import (
	"io"
	"regexp"

	ts "text/scanner"

	"toolman.org/numbers/stdsize"
	"toolman.org/timespan"
)

type Position = ts.Position

type Option interface {
	setOpt(s *Scanner)
}

type Scanner struct {
	mode     uint
	text     string
	token    int
	pos      Position
	timespan *timespan.Timespan
	regex    *regexp.Regexp
	stdSize  stdsize.Value
	doubles  map[rune]rune
	keywords map[string]int
	labels   map[rune]string
	pairs    map[rune]*RunePair
	gs       *ts.Scanner
}

// Source is a data source for Scanner text input. Any io.Reader having
// a `Name() string` method can be used as a Source, and -- as luck would
// have it, `*os.File` satisfies Source.
type Source interface {
	Name() string
	io.Reader
}

func New(src Source, options ...Option) *Scanner {
	gs := new(ts.Scanner).Init(src)

	gs.Mode = ts.GoTokens
	gs.Filename = src.Name()
	gs.Whitespace = defaultWhitespace

	s := &Scanner{
		gs:       gs,
		mode:     gs.Mode,
		doubles:  make(map[rune]rune),
		keywords: make(map[string]int),
		labels:   make(map[rune]string),
		pairs:    make(map[rune]*RunePair),
	}

	for _, o := range options {
		o.setOpt(s)
	}

	return s
}

func (s *Scanner) Scan() rune {
	tok := s.gs.Scan()
	s.text = s.gs.TokenText()
	s.pos = s.gs.Position
	s.timespan = nil
	s.regex = nil
	s.token = 0

	switch tok {
	case '#':
		if s.mode&uint(ScanHashComments) == 0 {
			return tok
		}
		return s.scanHashComment()

	case Ident:
		if t, ok := s.scanKeyword(); ok {
			return t
		}
		return tok

	case Int:
		if t := s.scanStdSize(tok); t != tok {
			return t
		}
		return s.scanTimespan(tok)

	case '/':
		if s.mode&uint(ScanRegexen) == 0 {
			return tok
		}
		return s.scanRegex()

	default:
		if ptok, ok := s.scanRunePair(tok); ok {
			return ptok
		}

		if dtok, ok := s.scanDoubles(tok); ok {
			return dtok
		}

		return tok
	}
}
func (s *Scanner) Next() rune   { return s.gs.Next() }
func (s *Scanner) Peek() rune   { return s.gs.Peek() }
func (s *Scanner) Text() string { return s.text }
func (s *Scanner) Token() int   { return s.token }

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

// Position returns the position of the most recently scanned token or, if that
// is invalid, the position of the character immediately following the most
// recently scanned token or character.
func (s *Scanner) Position() Position {
	switch {
	case s.pos.IsValid():
		return s.pos
	case s.gs.Position.IsValid():
		return s.gs.Position
	default:
		return s.gs.Pos()
	}
}
