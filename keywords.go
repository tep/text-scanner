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

/*
func KeywordMap(kwmap map[string]int) Option {
	kwords := make([]*Keyword, len(kwmap))
	var i int
	for w, t := range kwmap {
		kwords[i] = &Keyword{w, t}
		i++
	}
	return Keywords(kwords...)
}

func Keywords(kwords ...*Keyword) Option {
	return func(s *Scanner) {
		s.labels[KeyWord] = "Keyword"
		kwi := &keywordInfo{make(map[string]rune), make(map[rune]string)}
		for _, k := range kwords {
			r := rune(k.Token)
			kwi.byWord[k.Word] = r
			kwi.byRune[r] = k.Word
		}
		s.kwInfo = kwi
	}
}

func (s *Scanner) scanKeywords(tok rune) (rune, bool) {
	if s.kwInfo == nil {
		return tok, false
	}

	if kt, ok := s.kwInfo.byWord[s.text]; ok {
		s.keyword = &Keyword{s.text, int(kt)}
		return KeyWord, true
	}

	return tok, false
}

func (s *Scanner) Keyword() int {
	return s.keyword.Token
}
*/
