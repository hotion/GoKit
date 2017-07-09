package GoKit

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Message(t *testing.T) {
	ast := assert.New(t)

	ast.Equal("", Message(), "没有正确处理**空**消息")

	ast.Equal("foo", Message("foo"), "没能正确处理单个消息")

	ast.Equal("foobar", Message("foo%s", "bar"), "没能正确处理带参数的消息")

	ast.Equal("foobar, too", Message("foo%s, %s", "bar", "too"), "没能正确处理带两个参数的消息")
}
