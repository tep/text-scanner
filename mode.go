package scanner

type ScanMode int

func (o ScanMode) setOpt(s *Scanner) {
	if o < 0 {
		s.mode &^= uint(-o)
	} else {
		s.mode |= uint(o)
	}
}
