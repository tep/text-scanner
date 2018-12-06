package scanner

import (
	"fmt"
	"regexp"
	"strconv"
)

func init() { tokenStrings[Regex] = "Regex" }

func (s *Scanner) Regex() *regexp.Regexp {
	return s.regex
}

var qrPairs = map[rune]rune{
	'(': ')',
	'<': '>',
	'[': ']',
	'{': '}',
}

func init() {
	for _, r := range "!\"#$%&'*+,-./:;=?@^_`|~" {
		qrPairs[r] = r
	}
}

func (s *Scanner) canScanRegex(tok rune) (bool, rune) {
	if !s.can(ScanRegexen) {
		return false, 0
	}

	if tok == '/' {
		return true, '/'
	}

	if tok != 'r' {
		return false, 0
	}

	if end, ok := qrPairs[s.Peek()]; ok {
		s.Next()
		return true, end
	}

	return false, 0
}

func (s *Scanner) scanRegex(end rune) rune {
	sp := s.gs.Position

	res, nt, err := s.scanTo(end)
	if err != nil {
		s.error(err.Error())
		return nt
	}

	if flags := s.scanFlags(); flags != "" {
		res = fmt.Sprintf("(?%s:%s)", flags, res)
	}

	s.gs.Position = sp

	stringDump("regex", res)
	if s.regex, err = regexp.Compile(res); err != nil {
		s.error(fmt.Sprintf("%v: `%s`", err.Error(), res))
	}

	s.text = res
	return Regex
}

var stringDump = func(string, string) {}

func (s *Scanner) scanFlags() string {
	var flags string

	for done, nt := false, s.Peek(); !done; nt = s.Peek() {
		switch nt {
		case 'i', 'm', 's', 'U':
			flags += string(s.Next())

		default:
			done = true
		}
	}

	return flags
}

func (s *Scanner) scanTo(end rune) (string, rune, error) {
	var (
		out    string
		pt, nt rune
	)

	for {
		nt = s.Peek()

		if nt == EOF {
			return out, nt, fmt.Errorf("regex not terminated: `%s`", out)
		}

		if nt == end {
			if pt != '\\' {
				s.Next()
				break
			}

			out = out[:len(out)-1]
		}

		tok := s.Next()
		out += string(tok)

		if nt == '\\' && pt == '\\' {
			tok = 0
		}

		pt = tok
	}

	out, err := unquote(out)
	return out, nt, err
}

func unquote(str string) (string, error) {
	var out string

	for str != "" {
		v, _, t, err := strconv.UnquoteChar(str, 0)

		if err != nil {
			if err == strconv.ErrSyntax && str[0] == '\\' {
				str = "\\" + str
				continue
			}
			return out, err
		}

		out += string(v)
		str = t
	}

	return out, nil
}
