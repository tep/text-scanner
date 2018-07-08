package scanner

import ts "text/scanner"

const (
	// Token types exposed from "text/scanner"
	EOF       = ts.EOF
	Ident     = ts.Ident
	String    = ts.String
	RawString = ts.RawString
	Comment   = ts.Comment
	Float     = ts.Float
	Int       = ts.Int
)

const (
	// Custom token types
	customTokens = -(iota + 3 - ts.Comment)
	KeyWord      // A registered keyword
	HashComment  // A shell-style # comments
	Timespan     // A toolman.org/timespan literal
	Regex        // A regular expression literal
)

const (
	// Scan mode bits exposed from "text/scanner"
	ScanIdents     = ScanMode(ts.ScanIdents)
	ScanInts       = ScanMode(ts.ScanInts)
	ScanFloats     = ScanMode(ts.ScanFloats)
	ScanChars      = ScanMode(ts.ScanChars)
	ScanStrings    = ScanMode(ts.ScanStrings)
	ScanRawStrings = ScanMode(ts.ScanRawStrings)
	ScanComments   = ScanMode(ts.ScanComments)
	SkipComments   = ScanMode(ts.SkipComments)
	GoTokens       = ScanMode(ts.GoTokens)
)

const (
	// Custom mode bits
	customScans      = ScanMode(1 << -customTokens)
	ScanHashComments = ScanMode(1 << -HashComment)
	ScanTimespans    = ScanMode(1 << -Timespan)
	ScanRegexen      = ScanMode(1 << -Regex)
)

type ScanMode int

func (o ScanMode) setOpt(s *Scanner) {
	if o < 0 {
		s.mode &^= uint(-o)
	} else {
		s.mode |= uint(o)
	}
}
