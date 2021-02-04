package exception

import "fmt"

type ErrorCode struct {
	/* config-center 自身错误码 */
	code string
	/* http 错误码 */
	status int
	/* 错误信息 */
	message string
}

func (e * ErrorCode) GetCode() string {
	return e.code
}

func (e * ErrorCode) GetStatus() int {
	return e.status
}

func (e * ErrorCode) GetMessage() string {
	return e.message
}

func (e *ErrorCode) Error() string {
	return e.String()
}

func (e *ErrorCode) String() string {
	return fmt.Sprintf("{code:%s, status:%d, message:%s}", e.code, e.status, e.message)
}

func NewErrorCode(code string, status int, message string) *ErrorCode {
	return &ErrorCode{
		code:    code,
		status:  status,
		message: message,
	}
}

/* 0000-0099 公共 错误码用 */
var INTERVAL_ERROR = NewErrorCode("cc.0001", 500, "interval error")

var LOCK_EXIST = NewErrorCode("cc.0002", 500,"lock exist")

/* 0100-0199 给 management 用 */
var PROFILE_EXIST = NewErrorCode("cc.0100", 409, "profile is already existed")

/* 0100-0199 给 service 用 */
var PROFILE_NOT_EXIST = NewErrorCode("cc.0200", 404, "profile is not exist")