// Copyright © 2018 Timothy E. Peoples <eng@toolman.org>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package scanner

func (s *Scanner) Scan() (tok rune) {
	defer func() { s.reAfter.token(tok) }()
	switch tok = s.scanToken(); tok {
	case '#':
		return s.scanHashComment()

	case Ident:
		return s.scanIdent(tok)

	case Int:
		return s.scanInt(tok)

	default:
		return s.scanMisc(tok)
	}
}

func (s *Scanner) scanToken() rune {
	tok := s.gs.Scan()

	s.text = s.gs.TokenText()
	s.pos = s.gs.Position
	s.timespan = nil
	s.regex = nil
	s.token = int(tok)

	return tok
}

func (s *Scanner) scanIdent(tok rune) rune {
	if s.text == "r" {
		if ok, end := s.canScanRegex('r'); ok {
			return s.scanRegex(end)
		}
	}

	if t, ok := s.scanKeyword(); ok {
		return t
	}

	return tok
}

func (s *Scanner) scanInt(tok rune) rune {
	if t := s.scanStdSize(tok); t != tok {
		return t
	}
	return s.scanTimespan(tok)
}

// TODO(#1):  Add ability to tokenize whitespace following certain tokens.
//
//            Normally, whitespace is ignored and there are no whitespace
//            tokens returned by Scan().  However, temporarily emitting
//            whitespace tokens could prove quite useful.
//
//            We should add the ability for certain tokens to temporarily
//            enable the scanning of any immediately following whitespace
//            tokens, and then automatically disable it on the next non-
//            whitespace token.
//
//            If TOKEN_A is marked as a whitespace trigger token, the two
//            possibilities would be:
//
//							  TOKEN_A TOKEN_B
//
//							      This indicates that no whitespace whatsoever
//							      exists between these two tokens.
//
//							  TOKEN_A SPACE TOKEN_B
//
//							  		A SPACE token is emitted to indicate that one
//							  		or more whitespace characters (of any kind)
//							  		exist between the two tokens.
//
//            It is then up to the parser to deal with these SPACE tokens as
//            it sees fit.
//
//      NOTE: Instead of always emitting a single SPACE token regardless of
//            how much whitespace is scanned, we could:
//
//            		a) distinguish between types of whitespace
//            		   (space, tab, newline, etc... ) -- and/or
//
//            		b) emit a separate token for each and every
//            		   whitespace rune enountered.
//
//
func (s *Scanner) scanMisc(tok rune) rune {
	if ptok, ok := s.scanRunePair(tok); ok {
		return ptok
	}

	if dtok, ok := s.scanDoubles(tok); ok {
		return dtok
	}

	if ok, end := s.canScanRegex(tok); ok {
		return s.scanRegex(end)
	}

	return tok
}
