package stackerr

import "strings"
import "testing"
import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestWrap(t *testing.T) {
	err := fmt.Errorf("boom")
	err = Wrap(err)

	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("should contain the error message")
	}

	if !strings.Contains(err.Error(), "github.com") {
		t.Errorf("should contain the stack trace")
	}
}

func TestWrapNil(t *testing.T) {
	err := Wrap(nil)
	if err != nil {
		t.Error("should be nil")
	}
}
