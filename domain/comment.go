package domain

import "time"

// Comment this struct is comments model.
type Comment struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	PostID    int       `gorm:"primary_key" json:"post_id"`
	Sentence  string    `gorm:"colmun:sentence" json:"sentence"`
	Good      int       `gorm:"unsigned" json:"good"`
	IsEdited  bool      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Comments this type is slice from Comment struct.
type Comments []Comment
