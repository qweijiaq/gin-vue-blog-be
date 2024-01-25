package big_model

import (
	"fmt"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"log"
)

type QwenModel struct {
}

func (QwenModel) Send(content string) (reply string, err error) {
	accessKeyId := "LTAI5tFcbxzcWe2aEmAZzqh7"
	accessKeySecret := "h8QJEAzdBLeFZVO0HkN5mZdl9jYINM"
	agentKey := "002cba9a21164ab6870431ccec79c53b_p_efm"
	appId := "7e08b841ff174e4cbb155a1f9db0d1b3"

	// 尽量避免多次初始化
	tokenClient := client.AccessTokenClient{AccessKeyId: &accessKeyId, AccessKeySecret: &accessKeySecret, AgentKey: &agentKey}
	token, err := tokenClient.GetToken()
	if err != nil {
		log.Printf("%v\n", err)
		return "获取 token 失败", err
	}

	cc := client.CompletionClient{Token: &token}
	prompt := content

	request := &client.CompletionRequest{}
	request.SetAppId(appId)
	request.SetPrompt(prompt)

	response, err := cc.CreateCompletion(request)
	if err != nil {
		fmt.Println(err)
		return "生成回答失败", err
	}

	return response.GetData().GetText(), nil
}
