package models

import "time"

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	Created   time.Time  `gorm:"default:current_timestamp" json:"created"`
	Updated   time.Time  `gorm:"default:current_timestamp" json:"updated"`
	DeletedAt *time.Time `json:"deleted_at"`
}
