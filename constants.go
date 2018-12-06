package scanner

import ts "text/scanner"

const defaultWhitespace = ts.GoWhitespace | 1<<'\v' | 1<<'\f'

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
	StdSize      // A toolman.org/numbers/stdsize Value
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
	customScans = ScanMode(1 << -customTokens)

	// ScanHashComments is a scanner Option that enabled scanning of
	// hash comments. For Go style comments, see ScanComments. Similar
	// to ScanComments, the SkipComments option may be used to treat
	// comments as white space.
	ScanHashComments = ScanMode(1 << -HashComment)

	// ScanTimespans is a scanner Option that enabled scanning of
	// Timespan literals as defined by the toolman.org/timespan package.
	ScanTimespans = ScanMode(1 << -Timespan)

	// ScanRegexen is a scanner Option that enables unconditional
	// scanning of regular expression. For a more restrictive regular
	// expression option, which only enables scanning after specific
	// tokens, see ScanRegexenAfter.
	ScanRegexen = ScanMode(1 << -Regex)

	// ScanStdSizes is a scanner Option that enabled scanning ofstandard
	// size designations as defined by the toolman.org/numbers/stdsize
	// package.
	ScanStdSizes = ScanMode(1 << -StdSize)
)
