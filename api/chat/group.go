package chat

import (
	"encoding/json"
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/service/common/response"
	"server/utils"
	"strings"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

const (
	InRoomMsg  ctype.MsgType = 1 // 进入聊天室
	TextMsg    ctype.MsgType = 2 // 发文本消息
	ImageMsg   ctype.MsgType = 3 // 图片消息
	VoiceMsg   ctype.MsgType = 4 // 语音消息
	VideoMsg   ctype.MsgType = 5 // 视频消息
	SystemMsg  ctype.MsgType = 6 // 系统消息
	OutRoomMsg ctype.MsgType = 7 // 退出聊天室
)

type GroupRequest struct {
	Content string        `json:"content"`  // 聊天的内容
	MsgType ctype.MsgType `json:"msg_type"` // 聊天类型
}
type GroupResponse struct {
	NickName    string        `json:"nick_name"`    // 前端自己生成
	Avatar      string        `json:"avatar"`       // 头像
	MsgType     ctype.MsgType `json:"msg_type"`     // 聊天类型
	Content     string        `json:"content"`      // 聊天的内容
	OnlineCount int           `json:"online_count"` // 在线人数
	Date        time.Time     `json:"created_at"`   // 消息的时间
}

// ChatGroupView 群聊接口
// @Tags 聊天管理
// @Summary 群聊接口
// @Description 群聊接口 websocket
// @Param data body GroupRequest   false  "表示多个参数"
// @Router /api/chat_groups [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst)

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser
	// 需要去生成昵称，根据昵称首字关联头像地址
	// 昵称关联 addr
	logrus.Infof("%s %s 连接成功", addr, chatUser.NickName)
	SendMsg(addr, GroupResponse{
		NickName:    chatUser.NickName,
		Avatar:      chatUser.Avatar,
		MsgType:     SystemMsg,
		Content:     "进入聊天室",
		OnlineCount: len(ConnGroupMap),
		Date:        time.Now(),
	}, false)
	SendGroupMsg(conn, GroupResponse{
		NickName:    chatUser.NickName,
		Avatar:      chatUser.Avatar,
		Content:     fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
		Date:        time.Now(),
		OnlineCount: len(ConnGroupMap),
		MsgType:     InRoomMsg,
	})

	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     OutRoomMsg,
				Content:     fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap) - 1,
			})
			break
		}
		// 进行参数绑定
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "参数绑定失败",
				OnlineCount: len(ConnGroupMap),
			}, true)
			continue
		}
		// 判断类型
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupResponse{
					NickName:    chatUser.NickName,
					Avatar:      chatUser.Avatar,
					MsgType:     SystemMsg,
					Content:     "消息不能为空",
					OnlineCount: len(ConnGroupMap),
				}, false)
				continue
			}
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     request.Content,
				MsgType:     TextMsg,
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			SendGroupMsg(conn, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				Content:     fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:        time.Now(),
				OnlineCount: len(ConnGroupMap),
				MsgType:     InRoomMsg,
			})
		default:
			SendMsg(addr, GroupResponse{
				NickName:    chatUser.NickName,
				Avatar:      chatUser.Avatar,
				MsgType:     SystemMsg,
				Content:     "消息类型错误",
				OnlineCount: len(ConnGroupMap),
			}, true)
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(conn *websocket.Conn, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)

	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  true,
		MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}

// SendMsg 给某个用户发消息
func SendMsg(_addr string, response GroupResponse, isSave bool) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	if isSave {
		ip, addr := getIPAndAddr(_addr)
		global.DB.Create(&models.ChatModel{
			NickName: response.NickName,
			Avatar:   response.Avatar,
			Content:  response.Content,
			IP:       ip,
			Addr:     addr,
			IsGroup:  false,
			MsgType:  response.MsgType,
		})
	}
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	ip = addrList[0]
	addr = utils.GetAddr(ip)
	return ip, addr
}
