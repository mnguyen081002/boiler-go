package models

type Token struct {
	BaseModel
	UserID    string `json:"user_id" gorm:"column:user_id;type:uuid;not null;index:idx_user_id,unique"`
	ExpiredAt int64  `json:"expired_at" gorm:"column:expired_at;type:bigint;not null"`
}
