package GoKit

import "fmt"

// Message 把msgAndArgs 格式化成消息
// NOTICE: 使用方法 MessageFrom(msgAndArgs...)
func Message(msgAndArgs ...interface{}) string {
	switch len(msgAndArgs) {
	case 0:
		return ""
	case 1:
		return msgAndArgs[0].(string)
	default:
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}
