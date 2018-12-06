package scanner

type Double struct {
	Rune  rune
	Token int
}

func (d *Double) setOpt(s *Scanner) {
	t := rune(d.Token)
	s.doubles[d.Rune] = t
	s.labels[t] = "RuneDouble"
}

type Doubles []*Double

func (dd Doubles) setOpt(s *Scanner) {
	for _, d := range dd {
		d.setOpt(s)
	}
}

func (s *Scanner) scanDoubles(tok rune) (rune, bool) {
	if len(s.doubles) == 0 || s.Peek() != tok {
		return tok, false
	}

	var done bool
	if dt, ok := s.doubles[tok]; ok {
		s.text += string(s.Next())
		tok, done = dt, true
	}

	return tok, done
}
