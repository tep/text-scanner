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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

// If emitJS is true, we'll emit JSON for each encountered token instead of
// actually testing things. Use this to regenerate 'testdata/expected.json"
// when/if the test input changes significantly.
const emitJS = false

const (
	DOIT = iota + 2345
	BRACES
	ASSIGN
	LDANGLE
	RDANGLE
	L2BRACE
	R2BRACE
)

func TestScanner(t *testing.T) {
	defer func(sd func(string, string)) { stringDump = sd }(stringDump)

	stringDump = func(l, s string) {
		ss := make([]string, len(s))
		for i, c := range []byte(s) {
			ss[i] = strconv.QuoteRune(rune(c))
		}
		t.Logf("%s: [ %s ]", l, strings.Join(ss, " "))
	}

	wantList, err := loadExpected("testdata/expected.json")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}

	kwrds := KeywordMap{"doit": DOIT}

	pairs := RunePairs{
		{'{', '}', BRACES},
		{':', '=', ASSIGN},
	}

	dubls := Doubles{
		{'<', LDANGLE},
		{'>', RDANGLE},
		{'{', L2BRACE},
		{'}', R2BRACE},
	}

	var (
		tok rune
		i   int
		s   *Scanner
	)

	ef := ErrFunc(func(msg string) { err = fmt.Errorf("%s: %v", msg, s.Position()) })
	s = New(f, ef, kwrds, pairs, dubls, +ScanHashComments, +ScanTimespans, +ScanStdSizes, -SkipComments, +ScanRegexen)

	for tok != EOF {
		err = nil
		tok = s.Scan()
		want := wantList[i]

		got := expTok(s, tok, err)

		if emitJS {
			js, err := json.Marshal(got)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(string(js))
			// t.Logf("JSON: %s", js)
			continue
		}

		t.Logf("%d: %+v", i, got)
		if got.TokenString == "Regex" {
			t.Logf("    %s", showme(got.Text))
		}

		i++

		if !got.equal(want) {
			t.Errorf("token #%d=%d [pos=%v]: got %+v; wanted %+v", i, tok, s.Position(), got, want)
		}
	}
}

func expTok(s *Scanner, tok rune, err error) *expected {
	if err != nil {
		return &expected{Err: err.Error()}
	}

	var ts string
	switch tok {
	case KeyWord:
		ts = "Keyword"
	case BRACES:
		ts = "BRACES"
	case ASSIGN:
		ts = "ASSIGN"
	case LDANGLE:
		ts = "LDANGLE"
	case RDANGLE:
		ts = "RDANGLE"
	default:
		ts = s.TokenString(tok)
	}

	return &expected{TokenString: ts, Text: s.Text()}
}

//--------------------------------------------------------------

type expected struct {
	TokenString string `json:"token_string,omitempty"`
	Text        string `json:"text,omitempty"`
	Err         string `json:"error,omitempty"`
}

func loadExpected(filename string) ([]*expected, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var exp []*expected

	if err := json.Unmarshal(data, &exp); err != nil {
		return nil, err
	}

	return exp, nil
}

func (x *expected) equal(o *expected) bool {
	switch {
	case x == nil && o == nil:
		return true
	case x == nil || o == nil:
		return false
	case x.Err+o.Err != "" && strings.HasPrefix(x.Err, o.Err):
		return true
	case x.Err+o.Err != "" && strings.HasPrefix(o.Err, x.Err):
		return true
	case x.TokenString == o.TokenString && x.Text == o.Text:
		return true
	default:
		return false
	}
}

func showme(s string) string {
	var ss []string
	for _, r := range []rune(s) {
		ss = append(ss, fmt.Sprintf("'%c'", r))
	}

	return strings.Join(ss, ", ")
}
