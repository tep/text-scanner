package scanner

import (
	"toolman.org/time/timespan"
)

func init() { tokenStrings[Timespan] = "Timespan" }

func (s *Scanner) Timespan() *timespan.Timespan {
	return s.timespan
}

func (s *Scanner) scanTimespan(tok rune) rune {
	sp := s.gs.Position
	nt := s.Peek()

	if !s.can(ScanTimespans) || nt == EOF || !isTSLetter(nt) {
		return tok
	}

	for isTSChar(nt) {
		s.text += string(s.Next())
		nt = s.Peek()
	}

	s.gs.Position = sp
	var err error
	if s.timespan, err = timespan.ParseTimespan(s.text); err != nil {
		s.error(err.Error())
	}

	return Timespan
}

func isTSChar(r rune) bool {
	return isTSLetter(r) || ('0' <= r && r <= '9')
}

func isTSLetter(r rune) bool {
	switch r {
	case 'Y', 'M', 'W', 'D', 'd', 'h', 'm', 'n', 's', 'u', 'µ':
		return true
	}
	return false
}
