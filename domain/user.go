package domain

import "time"

// User this struct is user model.
type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    string    `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	UserName  string    `json:"user_name"`
	Img       string    `json:"user_img"`
	Pass      string    `gorm:"column:password" json:"pass"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Users this type is slice from User struct.
type Users []User
