package domain

import "time"

const TYPE_SHORT string = "short"
const TYPE_LONG string = "long"

// Post this struct is posts model.
type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	File      string    `json:"file"`
	Introduce string    `json:"introduce"`
	Type      string    `json:"type"`
	Good      int       `json:"good"`
	Watch     int       `json:"watch"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"-"`
}

// Posts this type is slice from Post struct.
type Posts []Post
