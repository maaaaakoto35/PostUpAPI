package domain

import "time"

// User this struct is user model.
type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	UserID    string    `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	UserName  string    `json:"user_name"`
	Img       string    `json:"user_img"`
	Pass      string    `gorm:"column:password" json:"pass"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Users this type is slice from User struct.
type Users []User

// ResUser this struct is for response user data.
type ResUser struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Img      string `json:"user_img"`
}

// ResUsers this type is slice from ResUser struct.
type ResUsers []ResUser

// BindUser this func is changing User into ResUser.
func BindUser(u User) ResUser {
	return ResUser{
		UserID:   u.UserID,
		UserName: u.UserName,
		Img:      u.Img,
	}
}

// BindUsers this func is changing Users into ResUsers.
func BindUsers(users Users) ResUsers {
	var newUsers ResUsers

	for _, u := range users {
		user := ResUser{
			UserID:   u.UserID,
			UserName: u.UserName,
			Img:      u.Img,
		}

		newUsers = append(newUsers, user)
	}

	return newUsers
}
