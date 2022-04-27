package models

import (
	"time"

	"gorm.io/gorm"
)

type ChatMembers struct {
	ID uint `json:"id" gorm:"primaryKey; autoIncrement;"`

	ChatId uint `json:"chat_id" gorm:"not null; index:idx_chat_user_deleted;"`

	UserId uint `json:"user_id" gorm:"not null; index:idx_chat_user_deleted;"`

	Type string `json:"type" gorm:"not null; type:enum('Admin', 'User')"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_chat_user_deleted;"`

	Chat *Chat `json:"chat"`
	User *User `json:"user"`
}
