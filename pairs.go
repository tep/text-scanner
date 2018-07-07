package scanner

type RunePair struct {
	Left  rune
	Right rune
	Token int
}

func Pairs(pairs ...*RunePair) Option {
	return func(s *Scanner) {
		m := make(map[rune]*RunePair)
		for _, p := range pairs {
			m[p.Left] = p
			s.labels[rune(p.Token)] = "RunePair"
		}
		s.pairs = m
	}
}

func (s *Scanner) scanPairs(tok rune) (rune, bool) {
	if s.pairs == nil {
		return tok, false
	}

	rp, ok := s.pairs[tok]

	if !ok {
		return tok, false
	}

	if rp.Right != s.Peek() {
		return tok, false
	}

	s.text += string(s.gs.Next())

	return rune(rp.Token), true
}
