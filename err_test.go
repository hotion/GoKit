package GoKit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Err(t *testing.T) {
	err := errors.New("bar")

	// 检查 Err
	e := Err(err, "foo")
	assert.Equal(t, e.Info, "foo", "Error.Info 不相等")
	assert.Equal(t, e.Prev, err, "Error.Prev 不相等")

	// 检查 输出
	e = Err(e, "more")
	expected := "more ==> foo ==> bar"
	actual := e.Error()

	assert.Equal(t, expected, actual, "格式输出出错")

	// 检查 第一个 error 是 *Error 的情况
	e = Err(nil, "bar")
	e = Err(e, "foo")
	e = Err(e, "more")

	expected = "more ==> foo ==> bar"
	actual = e.Error()

	assert.Equal(t, expected, actual, "格式输出出错")

}
