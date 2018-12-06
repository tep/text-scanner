package scanner

type Keyword struct {
	Word  string
	Token int
}

func (kw *Keyword) setOpt(s *Scanner) {
	s.labels[rune(kw.Token)] = "KeyWord"
	s.keywords[kw.Word] = kw.Token
}

type Keywords []*Keyword

func (kw Keywords) setOpt(s *Scanner) {
	for _, k := range kw {
		k.setOpt(s)
	}
}

type KeywordMap map[string]int

var xxx = KeywordMap{"foo": 2}

func (kwm KeywordMap) setOpt(s *Scanner) {
	for w, t := range kwm {
		s.keywords[w] = t
		s.labels[rune(t)] = "KeyWord"
	}
}

func (s *Scanner) scanKeyword() (tok rune, done bool) {
	if kt, ok := s.keywords[s.text]; ok {
		s.token = kt
		tok, done = KeyWord, true
	}
	return
}
