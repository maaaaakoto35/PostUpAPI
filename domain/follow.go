package domain

import "time"

// Follow this struct is follows model.
type Follow struct {
	ID              int       `gorm:"primary_key" json:"id"`
	FollowingUserID string    `gorm:"column:following_user_id" json:"following_user_id"`
	FollowedUserID  string    `gorm:"column:followed_user_id" json:"followed_user_id"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

// Follows this type is slice from Follow struct.
type Follows []Follow

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
			UserID: f.FollowedUserID,
		}
		newUsers = append(newUsers, follow)
	}

	return
}

// BindFolloweds func
func BindFolloweds(follows Follows) (newUsers ResUsers) {

	for _, f := range follows {
		follow := ResUser{
			UserID: f.FollowingUserID,
		}
		newUsers = append(newUsers, follow)
	}

	return
}
