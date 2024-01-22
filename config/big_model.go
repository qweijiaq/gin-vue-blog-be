package config

type Setting struct {
	Name      string `yaml:"name" json:"name"`
	Enable    bool   `yaml:"enable" json:"enable"`
	ApiKey    string `yaml:"api_key" json:"api_key"`
	ApiSecret string `yaml:"api_secret" json:"api_secret"`
	Title     string `yaml:"title" json:"title"`
	Prompt    string `yaml:"prompt" json:"prompt"`
}

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disable" json:"disable"`
}

type SessionSetting struct {
	ChatScope    int `yaml:"chat_scope" json:"chat_scope"`
	SessionScope int `yaml:"session_scope" json:"session_scope"`
	DayScope     int `yaml:"day_scope" json:"day_scope"`
}

type BigModel struct {
	Setting        Setting        `yaml:"setting"`         // 对话积分消耗
	ModelList      []ModelOption  `yaml:"model_list"`      // 会话的积分消耗
	SessionSetting SessionSetting `yaml:"session_setting"` // 每日赠送积分
}
