package entity

import (
	"time"

	_ "gorm.io/gorm" // 导入 gorm 包
)

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement"`                                             // 用户ID，主键，自动递增
	Username  string    `gorm:"size:50;unique;not null" json:"username"`                              // 用户名，唯一且不能为空
	Password  string    `gorm:"size:255;not null" json:"password"`                                    // 密码，不能为空（建议存储哈希值）
	Email     string    `gorm:"size:100;unique;not null" json:"email"`                                // 邮箱，唯一且不能为空
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"createdAt"`            // 创建时间，默认为当前时间
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdate" json:"updatedAt"` // 更新时间，默认为当前时间，每次更新时自动更新
}
