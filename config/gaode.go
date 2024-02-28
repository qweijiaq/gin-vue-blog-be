package config

type Gaode struct {
	Enable bool   `yaml:"enable" json:"enable"` // 是否启用
	Key    string `json:"key" yaml:"key"`       // key
}
