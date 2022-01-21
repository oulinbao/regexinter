// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package intersection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	type Case struct {
		Expr1  string
		Expr2  string
		Expect bool
	}
	cases := []Case{
		{"", "", true},
		{"a+", "a?", true},
		{"a*", "a+", true},
		{"a*", "a?", true},
		{"[a-zA-Z]+", "[a-z]+", true},
		{"a*", "b*", true},
		{"a*bba+", "b*aaab+a", true},
		{" ", `\s`, true},
		{"/api/v1/[0-9]+/get", `/api/v1/\w+/get`, true},
		//
		{"[A-Z]+", "[a-z]+", false},
		{"a", "b", false},
		{"a*bba+", "b*aaabbb+a", false},
		{"\\s+", "a+", false},
		{"/api/v1/.*/", "/api/v2/.*/", false},
		{"/api/v1/[0-9]+/get", "/api/v1/[a-z]+/get", false},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expect, HasIntersection(c.Expr1, c.Expr2))
	}
}
