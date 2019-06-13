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

type Keyword struct {
	Word  string
	Token int
}

func (kw *Keyword) setOpt(s *Scanner) {
	s.labels[rune(kw.Token)] = "KeyWord"
	s.keywords[kw.Word] = kw.Token
}

type Keywords []*Keyword

func (kw Keywords) setOpt(s *Scanner) {
	for _, k := range kw {
		k.setOpt(s)
	}
}

type KeywordMap map[string]int

var xxx = KeywordMap{"foo": 2}

func (kwm KeywordMap) setOpt(s *Scanner) {
	for w, t := range kwm {
		s.keywords[w] = t
		s.labels[rune(t)] = "KeyWord"
	}
}

func (s *Scanner) scanKeyword() (tok rune, done bool) {
	if kt, ok := s.keywords[s.text]; ok {
		s.token = kt
		tok, done = KeyWord, true
	}
	return
}
