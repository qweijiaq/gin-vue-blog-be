package big_model

import (
	"errors"
	"fmt"
	client "github.com/aliyun/alibabacloud-bailian-go-sdk/client"
	"log"
	"server/global"
	"server/models"
	"strconv"
)

type QwenModel struct {
	SessionID uint
}

func (qwen QwenModel) Send(content string) (reply interface{}, err error) {
	var accessKeyId = global.Config.BigModel.Setting.AccessKeyId
	var accessKeySecret = global.Config.BigModel.Setting.AccessKeySecret
	var agentKey = global.Config.BigModel.Setting.AgentKey
	var appId = global.Config.BigModel.Setting.AppId

	// 尽量避免多次初始化
	tokenClient := client.AccessTokenClient{AccessKeyId: &accessKeyId, AccessKeySecret: &accessKeySecret, AgentKey: &agentKey}
	token, err := tokenClient.GetToken()
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	cc := client.CompletionClient{Token: &token}
	prompt := content

	request := &client.CompletionRequest{}
	request.SetAppId(appId)
	request.SetStream(true)
	request.SetPrompt(prompt)

	sessionId := strconv.Itoa(int(qwen.SessionID))
	request.SetSessionId(sessionId)

	var sessionModel models.BigModelSessionModel
	err = global.DB.Preload("RoleModel").Take(&sessionModel, qwen.SessionID).Error
	if err != nil {
		return nil, errors.New("会话不存在")
	}
	rawPrompt := &client.ChatQaMessage{User: sessionModel.RoleModel.Prompt, Bot: "好的"}
	chatHistory := []*client.ChatQaMessage{rawPrompt}
	request.SetHistory(chatHistory)

	// 查当前这个会话都聊了哪些内容
	if qwen.SessionID != 0 {
		var chatList []models.BigModelChatModel
		global.DB.Find(&chatList, "session_id = ?", qwen.SessionID)
	}

	res, err := cc.CreateStreamCompletion(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}
