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

type Double struct {
	Rune  rune
	Token int
}

func (d *Double) setOpt(s *Scanner) {
	t := rune(d.Token)
	s.doubles[d.Rune] = t
	s.labels[t] = "RuneDouble"
}

type Doubles []*Double

func (dd Doubles) setOpt(s *Scanner) {
	for _, d := range dd {
		d.setOpt(s)
	}
}

func (s *Scanner) scanDoubles(tok rune) (rune, bool) {
	if len(s.doubles) == 0 || s.Peek() != tok {
		return tok, false
	}

	var done bool
	if dt, ok := s.doubles[tok]; ok {
		s.text += string(s.Next())
		tok, done = dt, true
	}

	return tok, done
}
