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
	// Scan mode bits exposed from "text/scanner"
	ScanIdents     = ts.ScanIdents
	ScanInts       = ts.ScanInts
	ScanFloats     = ts.ScanFloats
	ScanChars      = ts.ScanChars
	ScanStrings    = ts.ScanStrings
	ScanRawStrings = ts.ScanRawStrings
	ScanComments   = ts.ScanComments
	SkipComments   = ts.SkipComments
	GoTokens       = ts.GoTokens
)

const (
	// Custom token types
	customTokens = -(iota + 3 - ts.Comment)
	Keyword      // A registered keyword
	HashComment  // A shell-style # comments
	Timespan     // A toolman.org/timespan literal
	Regex        // A regular expression literal
)

const (
	// Custom mode bits
	customScans      = 1 << -customTokens
	ScanHashComments = 1 << -HashComment
	ScanTimespans    = 1 << -Timespan
	ScanRegexen      = 1 << -Regex
)
