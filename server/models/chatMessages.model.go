package models

import (
	"time"

	"gorm.io/gorm"
)

type ChatMessage struct {
	ID uint `json:"id" gorm:"primaryKey; autoIncrement;"`

	ChatId uint `json:"chat_id" gorm:"index:idx_chat_member_deleted;"`

	ChatMemberId uint `json:"chat_member_id" gorm:"index:idx_chat_member_deleted;"`

	Message string `json:"message" gorm:"type:tinyText; default:null;"`
	Media   string `json:"media" gorm:"size:150; default:null;"`

	Type string `json:"type" gorm:"column:message_type; not null; type:enum('Message', 'Media');"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAT time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_chat_member_deleted"`

	Chat *Chat `json:"chat"`

	ChatMember *ChatMembers `json:"chat_member"`
}
