package GoKit

import "testing"
import "os"
import "github.com/stretchr/testify/assert"

func Test_PWD(t *testing.T) {
	os.Chdir("/") // 改变目录为了适应不同的测试环境，比如，TravisCI
	expected := "/"
	assert.Equal(t, expected, PWD(), "获取了错误的PWD")
}

// FIXME: 修复没有权限的问题。
func Test_Mkdir(t *testing.T) {
	dir := "testdir"
	err := Mkdir(dir)
	assert.Nil(t, err, "创建%s文件夹时出错", dir)
	defer os.Remove(dir)

	assert.True(t, Exist(dir), "没有创建目录%s", dir)
}
