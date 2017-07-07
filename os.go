package GoKit

import "os"

// PWD like shell cmd pwd
// PWD 返回当前工作目录路径
// 只是忽略了可能发生的错误
func PWD() string {
	pwd, _ := os.Getwd()
	return pwd
}
