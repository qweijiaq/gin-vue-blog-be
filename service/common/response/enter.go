package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/utils"
	"server/utils/valid"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json() string {
	byteData, _ := json.Marshal(r)
	return string(byteData)
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}

func OkWithDataSSE(data any, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  "成功",
	}.Json()
	c.SSEvent("", content)
}

func OkWithSSE(data any, msg string, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  msg,
	}.Json()
	c.SSEvent("", content)
}

func OkWithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithMessageSSE(msg string, c *gin.Context) {
	data := Response{
		Code: Error,
		Data: map[string]any{},
		Msg:  msg,
	}.Json()
	c.SSEvent("", data)
}
func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, c)
}

// FailWithValidError 参数校验失败
func FailWithValidError(err error, c *gin.Context) {
	msg := valid.Error(err)
	FailWithMessage(msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
