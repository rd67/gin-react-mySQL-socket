package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	ID uint `json:"id" gorm:"primarykey;autoincrement;"`

	UserId uint `json:"user_id" gorm:"not null; index;"`

	AccessToken string `json:"access_token" gorm:"type:varchar(512); not null; index:idx_token_deleted;"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null; default:CURRENT_TIMESTAMP(3);"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null; default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index:idx_token_deleted;"`

	//	Relations

	User *User `json:"user"`
}
