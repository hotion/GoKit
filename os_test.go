package GoKit

import "testing"
import "os"
import "github.com/stretchr/testify/assert"

func Test_PWD(t *testing.T) {
	os.Chdir("/") // 改变目录为了适应不同的测试环境，比如，TravisCI
	expected := "/"
	assert.Equal(t, expected, PWD(), "获取了错误的PWD")
}
