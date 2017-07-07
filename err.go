// err.go
// Reference: https://github.com/reusee/codes/tree/master/err
//
//请牢记一点
//error是**接口**
//*Err实现Error()后，是满足error接口的类型
//

package GoKit

import "fmt"

// Error 是自定义的错误类型
type Error struct {
	Info string
	Prev error
}

// Err 对Error添加新的信息，以便于追踪错误。
func Err(information string, err error) *Error {
	return &Error{
		Info: information,
		Prev: err,
	}
}

// Error() 帮助 Error struct 实现 error 接口
func (e *Error) Error() string {
	if e.Prev == nil {
		// 只有当第一个 error 是 *Error时，才用的到这个
		return fmt.Sprintf("%s", e.Info)
	}
	return fmt.Sprintf("%s ==> %s", e.Info, e.Prev)
}
