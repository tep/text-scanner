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

type ScanMode int

func (o ScanMode) setOpt(s *Scanner) {
	if o < 0 {
		s.mode &^= uint(-o)
	} else {
		s.mode |= uint(o)
	}
}

func (s *Scanner) Enable(mode ...ScanMode) {
	for _, m := range mode {
		m.setOpt(s)
	}
}

func (s *Scanner) Disable(mode ...ScanMode) {
	for _, m := range mode {
		m *= -1
		m.setOpt(s)
	}
}

func (s *Scanner) can(m ScanMode) bool {
	return s.mode&uint(m) != 0
}
