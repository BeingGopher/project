package do

import "time"

type User struct {
	UserID    interface{} `gorm:"primaryKey;autoIncrement"`             // 用户ID，主键，自动递增
	Username  interface{} `gorm:"size:50;unique;not null"`              // 用户名，唯一且不能为空
	Password  interface{} `gorm:"size:255;not null"`                    // 密码，不能为空（建议存储哈希值）
	Email     interface{} `gorm:"size:100;unique;not null"`             // 邮箱，唯一且不能为空
	CreatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP"`            // 创建时间，默认为当前时间
	UpdatedAt time.Time   `gorm:"default:CURRENT_TIMESTAMP;autoUpdate"` // 更新时间，默认为当前时间，每次更新时自动更新
}
