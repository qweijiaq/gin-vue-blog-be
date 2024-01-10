package log_stash

import "server/global"

// NewRuntime 创建一个运行日志的log
func NewRuntime(serviceName string) *Action {
	log := &Action{
		serviceName: serviceName,
		logType:     RuntimeType,
	}
	var logModel LogModel
	err := global.DB.Take(&logModel,
		"type = ? and to_days(createdAt) = to_days(now()) and serviceName = ?", RuntimeType, serviceName).Error
	if err == nil {
		log.model = &logModel
	}
	return log
}
