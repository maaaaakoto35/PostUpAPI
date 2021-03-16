package database

import "github.com/maaaaakoto35/PostUpAPI/domain"

// FollowRepository this struct has SQLHandler.
type FollowRepository struct {
	SQLHandler
}

// FindConditions this func is finding follow by follow_id.
func (fr *FollowRepository) FindConditions(where ...interface{}) (follow domain.Follow, err error) {

	if err = fr.Find(&follow, where).Error; err != nil {
		return
	}
	return
}

// FindsByFollowing this func is finding follows by following.
func (fr *FollowRepository) FindsByFollowing(userID string) (follows domain.Follows, err error) {
	f := domain.Follow{
		FollowingUserID: userID,
	}
	if err = fr.Find(&follows, f).Error; err != nil {
		return
	}
	return
}

// FindsByFollowed this func is finding follows by followed.
func (fr *FollowRepository) FindsByFollowed(userID string) (follows domain.Follows, err error) {
	f := domain.Follow{
		FollowedUserID: userID,
	}
	if err = fr.Find(&follows, f).Error; err != nil {
		return
	}
	return
}

// Store this func is storing follow.
func (fr *FollowRepository) Store(f domain.Follow) (follow domain.Follow, err error) {
	if err = fr.Create(&f).Error; err != nil {
		return
	}
	follow = f
	return
}

// Update this func is updating follow.
func (fr *FollowRepository) Update(f domain.Follow) (follow domain.Follow, err error) {
	if err = fr.Save(&f).Error; err != nil {
		return
	}
	follow = f
	return
}

// UpdateValue this func is updating some columns in follow.
func (fr *FollowRepository) UpdateValue(f domain.Follow, set string, value string) (follow domain.Follow, err error) {
	if err = fr.SaveValue(&f, set, value).Error; err != nil {
		return
	}
	follow = f
	return
}

// DeleteByID this func is deletingfollow by id.
func (fr *FollowRepository) DeleteByID(follow domain.Follow) (err error) {
	if err = fr.Delete(&follow).Error; err != nil {
		return
	}
	return
}

// FindAll this func is finding all follows.
func (fr *FollowRepository) FindAll() (follows domain.Follows, err error) {
	if err = fr.Find(&follows).Error; err != nil {
		return
	}
	return
}
