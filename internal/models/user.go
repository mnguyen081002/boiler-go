package models

import (
	"erp/internal/constants"
	"time"
)

type User struct {
	BaseModel
	FirstName   string         `json:"first_name" gorm:"column:first_name;type:varchar(50);not null"`
	LastName    string         `json:"last_name" gorm:"column:last_name;type:varchar(50);not null"`
	Email       string         `json:"email" gorm:"column:email;type:varchar(100);not null;index:idx_email,unique"`
	Password    string         `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Social      string         `json:"social"`
	SocialID    string         `json:"social_id"`
	Role        constants.Role `json:"role" gorm:"column:role;type:varchar(50);not null"`
	LastLoginAt time.Time      `json:"last_login_at" gorm:"column:last_login_at;type:timestamp"`
}
