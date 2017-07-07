package GoKit

import "testing"
import "os"
import "github.com/stretchr/testify/assert"

func Test_PWD(t *testing.T) {
	os.Chdir("/")
	expected := "/"
	assert.Equal(t, expected, PWD(), "获取了错误的PWD")
}
