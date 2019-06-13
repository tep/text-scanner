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

type RunePair struct {
	Left  rune
	Right rune
	Token int
}

func (rp *RunePair) setOpt(s *Scanner) {
	s.pairs[rp.Left] = rp
	s.labels[rune(rp.Token)] = "RunePair"
}

type RunePairs []*RunePair

func (rp RunePairs) setOpt(s *Scanner) {
	for _, p := range rp {
		p.setOpt(s)
	}
}

func (s *Scanner) scanRunePair(tok rune) (rune, bool) {
	var done bool
	if p, ok := s.pairs[tok]; ok && s.Peek() == p.Right {
		s.text += string(s.Next())
		s.token = p.Token
		tok, done = rune(p.Token), true
	}
	return tok, done
}
