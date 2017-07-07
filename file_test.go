package GoKit

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func Test_Exist(t *testing.T) {
	ast := assert.New(t)

	filename := "Test_Exist.txt"
	os.Remove(filename)

	// 查看不存在的文件
	ast.True(!Exist(filename), "%s不存在，却反馈为存在", filename)

	os.Create(filename)
	defer os.Remove(filename)

	ast.True(Exist(filename), "%s存在，却反馈为不存在", filename)

}
