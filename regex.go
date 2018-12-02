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

func (s *Scanner) scanRegex() rune {
	sp := s.gs.Position

	res, nt, err := s.scanToSlash()
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
			flags += string(s.gs.Next())

		default:
			done = true
		}
	}

	return flags
}

func (s *Scanner) scanToSlash() (string, rune, error) {
	var (
		out    string
		pt, nt rune
	)

	for {
		nt = s.Peek()

		if nt == EOF {
			return out, nt, fmt.Errorf("regex not terminated: `%s`", out)
		}

		if nt == '/' {
			if pt != '\\' {
				s.gs.Next()
				break
			}

			out = out[:len(out)-1]
		}

		tok := s.gs.Next()
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
