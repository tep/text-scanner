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

import "toolman.org/numbers/stdsize"

func init() { tokenStrings[StdSize] = "StdSize" }

func (s *Scanner) StdSize() stdsize.Value { return s.stdSize }

func (s *Scanner) scanStdSize(tok rune) rune {
	sp := s.gs.Position
	nt := s.Peek()

	if !s.can(ScanStdSizes) || tok != Int || nt == EOF || !isSZLetter(nt) {
		return tok
	}

	s.text += string(s.Next())

	if t := s.Peek(); t == 'i' {
		s.text += string(s.Next())
	}

	var err error
	if s.stdSize, err = stdsize.Parse(s.text); err != nil {
		s.error(err.Error())
		return tok
	}

	s.gs.Position = sp
	return StdSize
}

func isSZLetter(r rune) bool {
	switch r {
	case 'K', 'M', 'G', 'T', 'P':
		return true
	default:
		return false
	}
}
