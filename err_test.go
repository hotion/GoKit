package GoKit

import (
	"errors"
	"testing"
)

func Test_NewErr(t *testing.T) {
	err := errors.New("haha")
	e := NewErr("hehe", err)
	t.Logf("%s", e)
}
