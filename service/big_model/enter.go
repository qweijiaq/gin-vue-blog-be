package big_model

import "errors"

type BigModelInterface interface {
	Send(content string) (reply string, err error)
}

func Send(name string, content string) (reply string, err error) {
	var ser BigModelInterface
	switch name {
	case "qwen":
		ser = QwenModel{}
	case "wenxin":
	case "xinghuo":
	case "tiangong":
	case "ChatGPT":
	default:
		return "暂不支持该大模型", errors.New("暂不支持的大模型")
	}
	return ser.Send(content)
}
