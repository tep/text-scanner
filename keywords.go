package scanner

type keywordInfo struct {
	byWord map[string]rune
	byRune map[rune]string
}

type KeyWord struct {
	Word  string
	Token int
}

func KeywordMap(kwmap map[string]int) Option {
	kwords := make([]*KeyWord, len(kwmap))
	var i int
	for w, t := range kwmap {
		kwords[i] = &KeyWord{w, t}
		i++
	}
	return Keywords(kwords...)
}

func Keywords(kwords ...*KeyWord) Option {
	return func(s *Scanner) {
		s.labels[Keyword] = "Keyword"
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
		s.keyword = &KeyWord{s.text, int(kt)}
		return Keyword, true
	}

	return tok, false
}

func (s *Scanner) Keyword() int {
	return s.keyword.Token
}
