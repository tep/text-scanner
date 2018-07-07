package scanner

type Double struct {
	Rune  rune
	Token int
}

func Doubles(dbls ...*Double) Option {
	return func(s *Scanner) {
		m := make(map[rune]rune)
		for _, d := range dbls {
			m[d.Rune] = rune(d.Token)
			s.labels[rune(d.Token)] = "RuneDouble"
		}
		s.doubles = m
	}
}

func (s *Scanner) scanDoubles(tok rune) (rune, bool) {
	if s.doubles == nil || s.Peek() != tok {
		return tok, false
	}

	if dt, ok := s.doubles[tok]; ok {
		s.text += string(s.gs.Next())
		return dt, true
	}

	return tok, false
}
