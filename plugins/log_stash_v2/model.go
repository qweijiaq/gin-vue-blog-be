package log_stash

import "time"

type LogModel struct {
	ID          uint      `gorm:"comment:id" json:"id" gorm:"primaryKey"`                         // 主键id
	CreatedAt   time.Time `gorm:"comment:添加时间" json:"created_at"`                                 // 添加时间
	UpdatedAt   time.Time `gorm:"comment:更新时间" json:"updated_at"`                                 // 更新时间
	IP          string    `gorm:"comment:ip" json:"ip"`                                           // ip
	Addr        string    `gorm:"comment:地址" json:"addr"`                                         // 地址
	Level       Level     `gorm:"comment:等级" json:"level"`                                        // 等级
	Title       string    `gorm:"comment:标题" json:"title"`                                        // 标题
	Content     string    `gorm:"comment:详情" json:"content"`                                      // 详情
	UserID      uint      `gorm:"comment:用户id" json:"userID"`                                     // 用户id
	UserName    string    `gorm:"comment:用户名" json:"userName"`                                    // 用户名
	ServiceName string    `gorm:"comment:服务名" json:"serviceName"`                                 // 服务名称
	Status      bool      `gorm:"comment:登录状态" json:"status"`                                     // 登录状态
	Type        LogType   `gorm:"comment:日志类型，1登录，2操作，3运行" json:"type"`                           // 日志的类型  1 登录 2 操作 3 运行
	ReadStatus  bool      `gorm:"column:readStatus;default:false;comment:阅读状态" json:"readStatus"` // 阅读状态   true   已读  false  未读
}
