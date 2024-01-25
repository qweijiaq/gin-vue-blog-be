package big_model

import (
	"errors"
	"server/global"
)

type BigModelInterface interface {
	Send(content string) (reply any, err error)
}

func Send(sessionId uint, content string) (reply any, err error) {
	var ser BigModelInterface
	switch global.Config.BigModel.Setting.Name {
	case "qwen":
		ser = QwenModel{SessionID: sessionId}
	case "wenxin":
	case "xinghuo":
	case "tiangong":
	case "ChatGPT":
	default:
		return "暂不支持该大模型", errors.New("暂不支持的大模型")
	}
	return ser.Send(content)
}
