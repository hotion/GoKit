package GoKit

import "testing"
import "github.com/stretchr/testify/assert"

func Test_IsSameType(t *testing.T) {
	inta, intb := 1, 2
	str := "3"

	assert.True(t, IsSameType(inta, intb), "把两个int类型的变量，判定为类型不一样")
	assert.False(t, IsSameType(inta, str), "把int型和string型变量，判定为类型一样")
}
