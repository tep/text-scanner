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

	if s.regex, err = regexp.Compile(res); err != nil {
		s.error(fmt.Sprintf("%v: `%s`", err.Error(), res))
	}

	s.text = res
	return Regex
}

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
	var esc bool
	var res string
	nt := s.Peek()

	// var toks []string
	// defer func() { fmt.Printf("TOKS: %v\n", toks) }()

	for done, nt := false, s.Peek(); !done; nt = s.Peek() {
		// toks = append(toks, strconv.QuoteRune(nt))

		switch nt {
		case '\n':
			if esc {
				esc = false
				continue
			}
			fallthrough

		case EOF:
			return res, nt, fmt.Errorf("regex not terminated: `%s`", res)

		case '\\':
			esc = !esc
			if esc {
				res += string(s.gs.Next())
				continue
			}

		case '/':
			if !esc {
				s.gs.Next()
				done = true
				continue
			}
			esc = false
		}

		tok := s.gs.Next()

		if esc {
			us, err := strconv.Unquote("\\" + string(tok))
			if err != nil {
				return res, nt, err
			}
			tok = rune(us[0])
			esc = false
		}
		res += string(tok)

		nt = s.Peek()
	}

	return res, nt, nil
}
