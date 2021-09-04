package domain

import "time"

const TYPE_SHORT string = "short"
const TYPE_LONG string = "long"

// Post this struct is posts model.
type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	UserName  string    `gorm:"-" json:"user_name"`
	UserImg   string    `gorm:"-" json:"user_img"`
	File      string    `json:"file"`
	Introduce string    `json:"introduce"`
	Type      string    `json:"type"`
	Good      int       `json:"good"`
	Watch     int       `json:"watch"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

// Posts this type is slice from Post struct.
type Posts []Post

// BindPost this func is changing User into ResUser.
func BindPost(p Post, r ResUser) Post {
	return Post{
		UserName:  p.UserID,
		UserImg:   r.Img,
		File:      p.File,
		Introduce: p.Introduce,
		Type:      p.Type,
		Good:      p.Good,
		Watch:     p.Watch,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// BindPosts this func is changing User into ResUsers.
func BindPosts(posts Posts, res map[string]ResUser) (results Posts) {
	for _, p := range posts {
		if r, ok := res[p.UserID]; ok {
			result := BindPost(p, r)
			results = append(results, result)
		}
	}
	return
}
