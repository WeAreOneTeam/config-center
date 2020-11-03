package code

import "fmt"

var (
	_codes = make(map[int]Code) // register codes.
)

type Code struct {
	ErrCode int
	ErrMsg  string
}

func New(e int, msg string) Code {
	if e <= 0 {
		panic("业务代码不能少于0")
	}
	return add(e, msg)
}

func add(e int, msg string) Code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("代码: %d 已存在", e))
	}
	code := new(Code)

	code.ErrMsg = msg
	_codes[e] = *code
	return *code
}

func (c Code) Message() string {
	if len(c.ErrMsg) > 0 {
		return c.ErrMsg
	}
	return _codes[c.ErrCode].ErrMsg
}
func (c Code) Code() int {
	return c.ErrCode
}

func (c Code) Error() string {
	if len(c.ErrMsg) > 0 {
		return c.ErrMsg
	}
	return _codes[c.ErrCode].ErrMsg
}

func (c Code) FormatMsgCode(msg string) Code {
	code := new(Code)
	code.ErrCode = c.ErrCode
	code.ErrMsg = fmt.Sprintf(c.ErrMsg, msg)
	return *code
}

func (c Code) MessageKey() string {
	key := fmt.Sprintf("backend.%d", c.ErrCode)
	return key
}
