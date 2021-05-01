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
	Short     Posts  `json:"short"`
	Long      Posts  `json:"long"`
}

// ResUsers this type is slice from ResUser struct.
type ResUsers []ResUser

type BindParam struct {
	ResUser   ResUser
	Following int
	Follower  int
	Short     *Posts
	Long      *Posts
}

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

func Bind(param BindParam) ResUser {
	if param.Short == nil {
		param.Short = &Posts{
			Post{},
		}
	}
	if param.Long == nil {
		param.Long = &Posts{
			Post{},
		}
	}
	return ResUser{
		UserID:    param.ResUser.UserID,
		UserName:  param.ResUser.UserName,
		Img:       param.ResUser.Img,
		Introduce: param.ResUser.Introduce,
		Following: param.Following,
		Follower:  param.Follower,
		Short:     *param.Short,
		Long:      *param.Long,
	}
}
