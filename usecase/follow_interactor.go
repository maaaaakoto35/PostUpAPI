package usecase

import (
	"github.com/maaaaakoto35/PostUpAPI/domain"
)

// FollowInteractor this struct has FollowRepository.
type FollowInteractor struct {
	FollowRepository FollowRepository
}

// FollowingUserID this func is from controller to repository.
func (fi *FollowInteractor) FollowingUserID(userID string) (users domain.ResUsers, err error) {
	follows, err := fi.FollowRepository.FindsByFollowing(userID)
	users = domain.BindFollowings(follows)
	return
}

// FollowingNum this func is a number of user's post.
func (fi *FollowInteractor) FollowingNum(userID string) (num int, err error) {
	follow := domain.Follow{
		FollowingUserID: userID,
	}
	num, err = fi.FollowRepository.CountConditions(follow)
	return
}

// FollowedUserID this func is from controller to repository.
func (fi *FollowInteractor) FollowedUserID(userID string) (users domain.ResUsers, err error) {
	follows, err := fi.FollowRepository.FindsByFollowed(userID)
	users = domain.BindFolloweds(follows)
	return
}

// FollowingNum this func is a number of user's post.
func (fi *FollowInteractor) FollowedNum(userID string) (num int, err error) {
	follow := domain.Follow{
		FollowedUserID: userID,
	}
	num, err = fi.FollowRepository.CountConditions(follow)
	return
}

// Add this func is from controller to repository.
func (fi *FollowInteractor) Add(f domain.Follow) (follow domain.Follow, err error) {
	follow, err = fi.FollowRepository.Store(f)
	return
}

// Update this func is from controller to repository.
func (fi *FollowInteractor) Update(f domain.Follow) (follow domain.Follow, err error) {
	follow, err = fi.FollowRepository.Update(f)
	return
}

// Delete this func is from controller to repository.
func (fi *FollowInteractor) Delete(f domain.Follow) (err error) {
	follow, err := fi.FollowRepository.FindConditions(f)
	err = fi.FollowRepository.DeleteByID(follow)
	return
}
