package config

type Setting struct {
	Name      string `yaml:"name"`
	Enable    bool   `yaml:"enable"`
	ApiKey    string `yaml:"api_key"`
	ApiSecret string `yaml:"api_secret"`
	Title     string `yaml:"title"`
	Prompt    string `yaml:"prompt"`
}

type ModelOption struct {
	Label   string `yaml:"label" json:"label"`
	Value   string `yaml:"value" json:"value"`
	Disable bool   `yaml:"disable" json:"disable"`
}

type BigModel struct {
	Setting   Setting       `yaml:"setting"`
	ModelList []ModelOption `yaml:"model_list"`
}
