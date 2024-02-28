package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"created_at"`     // 建站时间
	BeiAn       string `yaml:"bei_an" json:"bei_an"`             // 备案号
	Title       string `yaml:"title" json:"title"`               // 网站标题
	QQImage     string `yaml:"qq_image" json:"qq_image"`         // QQ 二维码
	Version     string `yaml:"version" json:"version"`           // 网站版本号
	Email       string `yaml:"email" json:"email"`               // 邮箱
	WechatImage string `yaml:"wechat_image" json:"wechat_image"` // 微信二维码
	Name        string `yaml:"name" json:"name"`                 // 作者昵称
	Job         string `yaml:"job" json:"job"`                   // 作者职业
	Addr        string `yaml:"addr" json:"addr"`                 // 作者地址
	Slogan      string `yaml:"slogan" json:"slogan"`             // 网站 Slogan
	SloganEn    string `yaml:"slogan_en" json:"slogan_en"`       // 网站英文 Slogan
	Web         string `yaml:"web" json:"web"`                   // 网站链接
	CSDNUrl     string `yaml:"csdn_url" json:"csdn_url"`         // CSDN 链接
	GiteeUrl    string `yaml:"gitee_url" json:"gitee_url"`       // Gitee 链接
	GithubUrl   string `yaml:"github_url" json:"github_url"`     // GitHub 链接
}
