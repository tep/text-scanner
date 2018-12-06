package scanner

type reAfter struct {
	prev int
	toks map[int]struct{}
}

// ScanRegexenAfter returns a scanner Option that enables the scanning of
// regular expressions immediately following one of the provided tokens.
// Regular expression scanning is otherwise disabled.
//
// This option should not be used in conjunction with the broader ScanRegexen,
// which enables regular expression scanning unconditionally.
func ScanRegexenAfter(tokens ...int) *reAfter {
	if len(tokens) == 0 {
		return nil
	}

	a := &reAfter{toks: make(map[int]struct{})}

	for _, t := range tokens {
		a.toks[t] = struct{}{}
	}

	return a
}

func (a *reAfter) setOpt(s *Scanner) {
	s.reAfter = a
}

func (a *reAfter) can() bool {
	if a == nil {
		return false
	}

	_, ok := a.toks[a.prev]
	return ok
}

func (a *reAfter) token(tok rune) {
	if a == nil {
		return
	}

	a.prev = int(tok)
}
