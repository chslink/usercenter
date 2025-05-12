package model

import (
	"time"
)

// UserStatus 用户状态枚举
type UserStatus int8

const (
	UserStatusActive UserStatus = 1
	UserStatusFrozen UserStatus = 2
)

// User 用户核心表结构
type User struct {
	UID          string     `json:"uid" db:"uid"`
	Username     string     `json:"username" db:"username"`
	Phone        string     `json:"phone" db:"phone"`
	Email        *string    `json:"email,omitempty" db:"email"`
	PasswordHash string     `json:"-" db:"password_hash"`
	Status       UserStatus `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}
