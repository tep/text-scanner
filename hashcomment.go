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

func init() { tokenStrings[HashComment] = "HashComment" }

func (s *Scanner) scanHashComment() rune {
	if !s.can(ScanHashComments) {
		return rune(s.token)
	}

	sp := s.gs.Position
	cs := s.scanToEOL()

	if s.can(SkipComments) {
		// If we're skipping comments, we'll need to advance the scanner...
		s.Next()

		// reset the scanner's Position...
		s.gs.Position = s.gs.Pos()

		// then return the next scanned token
		return s.Scan()
	}

	s.text = "#" + cs
	s.gs.Position = sp

	return HashComment
}

func (s *Scanner) scanToEOL() string {
	var cs string
	nt := s.Peek()

	for nt != EOF && nt != '\n' {
		cs += string(s.Next())
		nt = s.Peek()
	}

	return cs
}
