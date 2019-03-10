// Licensed under the MIT license, see LICENCE file for details.

package quicktest_test

import (
	"errors"
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestBadCheckf(t *testing.T) {
	err := qt.BadCheckf("bad %s", "wolf")
	expectedMessage := "bad check: bad wolf"
	if err.Error() != expectedMessage {
		t.Fatalf("error:\ngot  %q\nwant %q", err, expectedMessage)
	}
}

func TestIsBadCheck(t *testing.T) {
	err := qt.BadCheckf("bad wolf")
	assertBool(t, qt.IsBadCheck(err), true)
	err = errors.New("bad wolf")
	assertBool(t, qt.IsBadCheck(err), false)
}

var errBadWolf = &errTest{}

// errTest is an error type used in tests.
type errTest struct{}

// Error implements error.
func (*errTest) Error() string {
	return "bad wolf"
}

// Format implements fmt.Formatter.
func (err *errTest) Format(f fmt.State, c rune) {
	if !f.Flag('+') || c != 'v' {
		fmt.Fprint(f, "unexpected verb for formatting the error")
	}
	fmt.Fprint(f, err.Error())
	fmt.Fprint(f, "\n  file:line")
}
