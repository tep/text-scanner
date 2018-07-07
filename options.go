package scanner

import ts "text/scanner"

type Option func(s *Scanner)

func Error(errFunc func(string)) Option {
	return func(s *Scanner) {
		s.gs.Error = func(_ *ts.Scanner, msg string) {
			errFunc(msg)
		}
	}
}

// Modes returns an Option to adjust which tokens the Scanner will return based
// on the provided modes. Each mode should be a value appropriate for the Mode
// field of text/scanner.Scanner or one of the custom scan modes defined in
// this package (i.e. ScanHashComments, ScanTimespans or ScanRegexen).
//
// Positive values will enable the given mode; negative values will disable it.
//
// To illistrate, we'll modify Scanner's default scan mode (text/scanner's
// GoTokens - which includes SkipComments). To disable SkipComments and also
// enable scanning of HashComments, use the following Option:
//
// 		scanner.Modes(+scanner.ScanHashComments, -scanner.SkipComments)
//
func Modes(modes ...int) Option {
	return func(s *Scanner) {
		for _, m := range modes {
			fmod, v := modifierFor(m)
			for _, fld := range s.fieldsFor(v) {
				fmod(fld, v)
			}
		}
	}
}

func (s *Scanner) fieldsFor(m uint) []*uint {
	if m > customScans {
		return []*uint{&s.mode}
	}
	return []*uint{&s.mode, &s.gs.Mode}
}

type fieldModifier func(*uint, uint)

func modifierFor(mode int) (fieldModifier, uint) {
	fm := func(fp *uint, v uint) { *fp |= v }
	if mode < 0 {
		mode = -mode
		fm = func(fp *uint, v uint) { *fp &^= v }
	}

	return fm, uint(mode)
}
