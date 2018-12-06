package scanner

import "toolman.org/numbers/stdsize"

func init() { tokenStrings[StdSize] = "StdSize" }

func (s *Scanner) StdSize() stdsize.Value { return s.stdSize }

func (s *Scanner) scanStdSize(tok rune) rune {
	sp := s.gs.Position
	nt := s.Peek()

	if s.mode&uint(ScanStdSizes) == 0 || tok != Int || nt == EOF || !isSZLetter(nt) {
		return tok
	}

	s.text += string(s.Next())

	if t := s.Peek(); t == 'i' {
		s.text += string(s.Next())
	}

	var err error
	if s.stdSize, err = stdsize.Parse(s.text); err != nil {
		s.error(err.Error())
		return tok
	}

	s.gs.Position = sp
	return StdSize
}

func isSZLetter(r rune) bool {
	switch r {
	case 'K', 'M', 'G', 'T', 'P':
		return true
	default:
		return false
	}
}
