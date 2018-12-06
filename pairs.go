package scanner

type RunePair struct {
	Left  rune
	Right rune
	Token int
}

func (rp *RunePair) setOpt(s *Scanner) {
	s.pairs[rp.Left] = rp
	s.labels[rune(rp.Token)] = "RunePair"
}

type RunePairs []*RunePair

func (rp RunePairs) setOpt(s *Scanner) {
	for _, p := range rp {
		p.setOpt(s)
	}
}

func (s *Scanner) scanRunePair(tok rune) (rune, bool) {
	var done bool
	if p, ok := s.pairs[tok]; ok && s.Peek() == p.Right {
		s.text += string(s.Next())
		s.token = p.Token
		tok, done = rune(p.Token), true
	}
	return tok, done
}
