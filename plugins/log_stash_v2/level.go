package log_stash

import "encoding/json"

type Level int

const (
	Info    Level = 1
	Warning Level = 2
	Error   Level = 3
)

// String 转字符串
func (level Level) String() string {
	switch level {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	}
	return ""
}

// MarshalJSON 自定义类型转换为json
func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}
