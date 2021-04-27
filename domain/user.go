package domain

import "time"

// User this struct is user model.
type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"unique" json:"user_id"`
	UserName  string    `json:"user_name"`
	Img       string    `json:"user_img"`
	Introduce string    `json:"introduce"`
	Pass      string    `gorm:"column:password" json:"pass"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Users this type is slice from User struct.
type Users []User

// ResUser this struct is for response user data.
type ResUser struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Img       string `json:"user_img"`
	Introduce string `json:"introduce"`
	Follower  int    `json:"follower"`
	Following int    `json:"following"`
}

// ResUsers this type is slice from ResUser struct.
type ResUsers []ResUser

// BindUser this func is changing User into ResUser.
func BindUser(u User) ResUser {
	return ResUser{
		UserID:    u.UserID,
		UserName:  u.UserName,
		Img:       u.Img,
		Introduce: u.Introduce,
	}
}

// BindUsers this func is changing Users into ResUsers.
func BindUsers(users Users) ResUsers {
	var newUsers ResUsers

	for _, u := range users {
		user := ResUser{
			UserID:    u.UserID,
			UserName:  u.UserName,
			Img:       u.Img,
			Introduce: u.Introduce,
		}

		newUsers = append(newUsers, user)
	}

	return newUsers
}

func BindFF(resUser ResUser, following int, follower int) ResUser {
	return ResUser{
		UserID:    resUser.UserID,
		UserName:  resUser.UserName,
		Img:       resUser.Img,
		Introduce: resUser.Introduce,
		Following: following,
		Follower:  follower,
	}
}
