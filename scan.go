package scanner

func (s *Scanner) Scan() (tok rune) {
	switch tok = s.scanToken(); tok {
	case '#':
		return s.scanHashComment()

	case Ident:
		return s.scanIdent(tok)

	case Int:
		return s.scanInt(tok)

	default:
		return s.scanMisc(tok)
	}
}

func (s *Scanner) scanToken() rune {
	tok := s.gs.Scan()

	s.text = s.gs.TokenText()
	s.pos = s.gs.Position
	s.timespan = nil
	s.regex = nil
	s.token = int(tok)

	return tok
}

func (s *Scanner) scanIdent(tok rune) rune {
	if t, ok := s.scanKeyword(); ok {
		return t
	}

	return tok
}

func (s *Scanner) scanInt(tok rune) rune {
	if t := s.scanStdSize(tok); t != tok {
		return t
	}
	return s.scanTimespan(tok)
}

func (s *Scanner) scanMisc(tok rune) rune {
	if ptok, ok := s.scanRunePair(tok); ok {
		return ptok
	}

	if dtok, ok := s.scanDoubles(tok); ok {
		return dtok
	}

	if ok, end := s.canScanRegex(tok); ok {
		return s.scanRegex(end)
	}

	return tok
}
