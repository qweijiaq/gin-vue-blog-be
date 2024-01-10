package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log_stash "server/plugins/log_stash_v2"
)

type responseWrite struct {
	gin.ResponseWriter
	byteData *bytes.Buffer
}

func (rw responseWrite) Write(buf []byte) (int, error) {
	rw.byteData.Write(buf)
	return rw.ResponseWriter.Write(buf)
}

func LogMiddleWare() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		// 请求
		r := responseWrite{
			ResponseWriter: c.Writer,
			byteData:       bytes.NewBuffer([]byte{}),
		}
		c.Writer = r
		c.Next()
		// 响应
		_action, ok := c.Get("action")
		if !ok {
			return
		}
		action, ok := _action.(*log_stash.Action)
		if !ok {
			return
		}
		action.SetResponseContent(r.byteData.String())
		action.SetFlush()
	}
}
