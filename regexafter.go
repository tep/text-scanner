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

type reAfter struct {
	prev int
	toks map[int]struct{}
}

// ScanRegexenAfter returns a scanner Option that enables the scanning of
// regular expressions immediately following one of the provided tokens.
// Regular expression scanning is otherwise disabled.
//
// This option should not be used in conjunction with the broader ScanRegexen,
// which enables regular expression scanning unconditionally.
func ScanRegexenAfter(tokens ...int) *reAfter {
	if len(tokens) == 0 {
		return nil
	}

	a := &reAfter{toks: make(map[int]struct{})}

	for _, t := range tokens {
		a.toks[t] = struct{}{}
	}

	return a
}

func (a *reAfter) setOpt(s *Scanner) {
	s.reAfter = a
}

func (a *reAfter) can() bool {
	if a == nil {
		return false
	}

	_, ok := a.toks[a.prev]
	return ok
}

func (a *reAfter) token(tok rune) {
	if a == nil {
		return
	}

	a.prev = int(tok)
}
