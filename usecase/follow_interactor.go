package usecase

import (
	"github.com/maaaaakoto35/PostUpAPI/domain"
)

// FollowInteractor this struct has FollowRepository.
type FollowInteractor struct {
	FollowRepository FollowRepository
}

// FollowingUserID this func is from controller to repository.
func (fi *FollowInteractor) FollowingUserID(userID string) (user domain.ResUser, err error) {
	f := domain.Follow{
		FollowingUserID: userID,
	}
	follow, err := fi.FollowRepository.FindConditions(f)
	user = domain.BindFollowing(follow)
	return
}

// FollowedUserID this func is from controller to repository.
func (fi *FollowInteractor) FollowedUserID(userID string) (user domain.ResUser, err error) {
	f := domain.Follow{
		FollowedUserID: userID,
	}
	follow, err := fi.FollowRepository.FindConditions(f)
	user = domain.BindFollowed(follow)
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
func (fi *FollowInteractor) Delete(followingUserID, followedUserID string) (err error) {
	f := domain.Follow{
		FollowingUserID: followingUserID,
		FollowedUserID:  followedUserID,
	}
	follow, err := fi.FollowRepository.FindConditions(f)
	err = fi.FollowRepository.DeleteByID(follow)
	return
}
