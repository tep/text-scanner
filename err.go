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
	"fmt"
	"os"

	ts "text/scanner"
)

func (s *Scanner) error(msg string) {
	if s.gs.Error != nil {
		s.gs.Error(s.gs, msg)
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %s\n", s.Position(), msg)
}

type ErrFunc func(string)

func (ef ErrFunc) setOpt(s *Scanner) {
	s.gs.Error = func(_ *ts.Scanner, msg string) { ef(msg) }
}
