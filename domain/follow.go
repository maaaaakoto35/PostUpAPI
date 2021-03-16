package domain

import "time"

// Follow this struct is posts model.
type Follow struct {
	FollowingUserID string    `gorm:"column:following_user_id" json:"following_user_id"`
	FollowedUserID  string    `gorm:"column:followed_user_id" json:"followed_user_id"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

// Follows this type is slice from Post struct.
type Follows []Follow

// TableName overrides the table name used by User to `profiles`
func (Follow) TableName() string {
	return "follows"
}

// BindFollowing this func is changing User into ResUser.
func BindFollowing(f Follow) ResUser {
	return ResUser{
		UserID: f.FollowingUserID,
	}
}

// BindFollowed this func is changing User into ResUser.
func BindFollowed(f Follow) ResUser {
	return ResUser{
		UserID: f.FollowedUserID,
	}
}

// BindFollowings func
func BindFollowings(follows Follows) (newUsers ResUsers) {

	for _, f := range follows {
		follow := ResUser{
			UserID: f.FollowingUserID,
		}
		newUsers = append(newUsers, follow)
	}

	return
}

// BindFolloweds func
func BindFolloweds(follows Follows) (newUsers ResUsers) {

	for _, f := range follows {
		follow := ResUser{
			UserID: f.FollowedUserID,
		}
		newUsers = append(newUsers, follow)
	}

	return
}
