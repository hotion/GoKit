package GoKit

import "testing"
import "os"
import "github.com/stretchr/testify/assert"

func Test_PWD(t *testing.T) {
	oldPWD := PWD()
	defer os.Chdir(oldPWD)

	os.Chdir("/") // 改变目录为了适应不同的测试环境，比如，TravisCI
	expected := "/"
	assert.Equal(t, expected, PWD(), "获取了错误的PWD")
}

func Test_Mkdir(t *testing.T) {
	dir := PWD() + "testdir"
	err := Mkdir(dir)
	assert.Nil(t, err, `创建"%s"文件夹时出错`, dir)
	defer os.Remove(dir)
	assert.True(t, Exist(dir), "没有创建目录%s", dir)

	dir = PWD() + "/noExist/testdir"
	err = Mkdir(dir)
	assert.NotNil(t, err, `创建"%s"文件夹时出错`, dir)
	assert.False(t, Exist(dir), "没有创建目录%s", dir)
}
