package domain

import "time"

// Post this struct is posts model.
type Post struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	UserID    string `json:"user_id"`
	File      string
	Introduce string
	Good      string
	Watch     string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Posts this type is slice from Post struct.
type Posts []Post
