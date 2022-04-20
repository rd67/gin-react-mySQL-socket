package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primarykey;autoIncrement;"`

	Name string `json:"name" gorm:"size:100;not null;"`

	Email string `json:"email" gorm:"size:100;not null;index:idx_email_deleted;" binding:"required"`

	Password string `json:"password" gorm:"type:tinytext; not null;"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_email_deleted;"`

	//	Relations
	UserTokens []UserToken `json:"user_tokens" gorm:"foreignKey:user_id;"`
	// AccessToken  UserToken   `json:"user_tokens" gorm:"foreignKey:user_id;"`
}
