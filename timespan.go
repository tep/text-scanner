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

import (
	"toolman.org/time/timespan/v2"
)

func init() { tokenStrings[Timespan] = "Timespan" }

func (s *Scanner) Timespan() *timespan.Timespan {
	return s.timespan
}

func (s *Scanner) scanTimespan(tok rune) rune {
	sp := s.gs.Position
	nt := s.Peek()

	if !s.can(ScanTimespans) || nt == EOF || !isTSLetter(nt) {
		return tok
	}

	for isTSChar(nt) {
		s.text += string(s.Next())
		nt = s.Peek()
	}

	s.gs.Position = sp
	var err error
	if s.timespan, err = timespan.ParseTimespan(s.text); err != nil {
		s.error(err.Error())
	}

	return Timespan
}

func isTSChar(r rune) bool {
	return isTSLetter(r) || ('0' <= r && r <= '9')
}

func isTSLetter(r rune) bool {
	switch r {
	case 'Y', 'M', 'W', 'D', 'd', 'h', 'm', 'n', 's', 'u', 'µ':
		return true
	}
	return false
}
