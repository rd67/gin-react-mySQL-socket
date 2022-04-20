package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID uint `json:"id" gorm:"primaryKey; autoIncrement;"`

	UserId uint `json:"user_id" gorm:"not null; index:idx_user_ouser_deleted;"`

	OUserId uint `json:"ouser_id" gorm:"column:ouser_id; not null; index:idx_user_ouser_deleted;"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_user_ouser_deleted"`

	// //	Relations
	User  *User `json:"user"`
	OUser *User `json:"ouser"`

	ChatMessages *[]ChatMessage `json:"chat_messages"`
}
