package scanner

func init() { tokenStrings[HashComment] = "HashComment" }

func (s *Scanner) scanHashComment() rune {
	if !s.can(ScanHashComments) {
		return rune(s.token)
	}

	sp := s.gs.Position
	cs := s.scanToEOL()

	if s.can(SkipComments) {
		// If we're skipping comments, we'll need to advance the scanner...
		s.Next()

		// reset the scanner's Position...
		s.gs.Position = s.gs.Pos()

		// then return the next scanned token
		return s.Scan()
	}

	s.text = "#" + cs
	s.gs.Position = sp

	return HashComment
}

func (s *Scanner) scanToEOL() string {
	var cs string
	nt := s.Peek()

	for nt != EOF && nt != '\n' {
		cs += string(s.Next())
		nt = s.Peek()
	}

	return cs
}
