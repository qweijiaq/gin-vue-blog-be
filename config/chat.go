package config

type ChatGroup struct {
	IsAnonymous    bool   `yaml:"isAnonymous" json:"isAnonymous"`       // 匿名群聊
	IsShowTime     bool   `yaml:"isShowTime" json:"isShowTime"`         // 显示时间
	DefaultLimit   int    `yaml:"defaultLimit" json:"defaultLimit"`     // 初识条数
	ContentLength  int    `yaml:"contentLength" json:"contentLength"`   // 初识条数
	WelcomeTitle   string `yaml:"welcomeTitle" json:"welcomeTitle"`     // 欢迎语
	IsOnlinePeople bool   `yaml:"isOnlinePeople" json:"isOnlinePeople"` // 显示人数
	IsSendImage    bool   `yaml:"isSendImage" json:"isSendImage"`       // 可发图片
	IsSendFile     bool   `yaml:"isSendFile" json:"isSendFile"`         // 可发文件
	IsMd           bool   `yaml:"isMd" json:"isMd"`                     // 是否支持md
}
