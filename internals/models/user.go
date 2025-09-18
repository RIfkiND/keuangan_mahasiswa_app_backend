package models

import "time"

type User struct {
    ID       uint      `json:"id" gorm:"primaryKey"`
    Name     string    `json:"name"`
    Email    string    `json:"email" gorm:"uniqueIndex"`
    Password string    `json:"password"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
    Post      []Post
}								