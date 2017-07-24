package GoKit

import (
	"fmt"
	"os"
)

// PWD like shell cmd pwd
// PWD 返回当前工作目录路径
// 只是忽略了可能发生的错误
func PWD() string {
	pwd, _ := os.Getwd()
	return pwd
}

// Mkdir 简化了os.Mkdir的操作
// 会按照name，新建一个权限为0755的目录
func Mkdir(name string) error {
	// mask := syscall.Umask(0)
	// defer syscall.Umask(mask)

	err := os.Mkdir(name, 0755)
	if err != nil {
		return fmt.Errorf("无法创建文件夹%s: %s", name, err)
	}

	return nil
}
