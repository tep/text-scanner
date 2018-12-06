package scanner

type ScanMode int

func (o ScanMode) setOpt(s *Scanner) {
	if o < 0 {
		s.mode &^= uint(-o)
	} else {
		s.mode |= uint(o)
	}
}
func (s *Scanner) can(m ScanMode) bool {
	return s.mode&uint(m) != 0
}
