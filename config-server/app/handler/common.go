package handler

import (
	"config-center/config-server/app/consts/code"
	"config-center/config-server/app/model/response"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	ecode := code.OK
	resp := &response.BaseResponse{
		Status: ecode.Code(),
		Msg:    ecode.Message(),
		Data:   data,
	}
	ctx.PureJSON(200, resp)
}

func FailWithError(c *gin.Context, err error) {
	ecode, ok := err.(code.Code)
	if ok {
		Fail(c, ecode)
	} else {
		Fail(c, code.UnknownError)
	}
}

func Fail(ctx *gin.Context, ecode code.Code) {
	resp := &response.BaseResponse{
		Status: ecode.Code(),
		Msg:    ecode.MessageKey(),
		Data:   make(map[string]string),
	}
	ctx.JSON(200, resp)
}
