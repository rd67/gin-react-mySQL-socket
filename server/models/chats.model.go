package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID uint `json:"id" gorm:"primaryKey; autoIncrement;"`

	Name string `json:"name" gorm:"size:255; default: null;" `

	Type string `json:"type" gorm:"not null; type:enum('Single', 'Group');"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_deleted"`

	ChatMembers *[]ChatMembers `json:"chat_members"`

	ChatMessages *[]ChatMessage `json:"chat_messages"`
}
