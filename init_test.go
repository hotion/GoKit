package GoKit

import (
	"os"
	"runtime"
	"testing"

	"syscall"

	"io/ioutil"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func Test_Init(t *testing.T) {
	filename := "test.pid"

	Init(filename)
	defer os.Remove(filename)

	// 检查核心数设置
	expected := runtime.GOMAXPROCS(-1)
	actual := runtime.NumCPU()
	assert.Equal(t, expected, actual, "没能把GOMAXPROCS设置成最大核心数")

	// 检查是否正确保存了pid
	expectedPID := fmt.Sprint(syscall.Getpid())
	data, _ := ioutil.ReadFile(filename)
	actualPID := string(data)
	assert.Equal(t, expectedPID, actualPID, "保存的PID的实际的不一致。")
}

func Test_WaitingKill(t *testing.T) {
	ast := assert.New(t)
	done := WaitingKill()

	go func() {
		// 获取当前进程的 pid
		pid := syscall.Getpid()

		// 通过pid号，找到当前进程
		p, err := os.FindProcess(pid)
		ast.Nil(err, "CANNOT find %d Procesee", pid)

		// 给当前进程发送关闭信号 syscall.SIGINT
		// 相当于在命令行，按下了 Ctrl+c
		err = p.Signal(syscall.SIGINT)
		ast.Nil(err, "FAIL to send syscall.SIGINT to %d Process", pid)
	}()

	data, ok := <-done
	assert.False(t, ok, "从关闭通道done中，接收到了实际传输过来的数据: %v", data)
}
