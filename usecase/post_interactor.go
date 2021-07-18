package usecase

import (
	"sort"

	"github.com/maaaaakoto35/PostUpAPI/domain"
)

// PostInteractor this struct has PostRepository.
type PostInteractor struct {
	PostRepository PostRepository
}

// PostByID this func is from controller to repository.
func (pi *PostInteractor) PostByID(id int) (post domain.Post, err error) {
	post, err = pi.PostRepository.FindByID(id)
	return
}

// PostsByUserID this func is from controller to repository.
func (pi *PostInteractor) PostsByUserID(userID string) (posts domain.Posts, err error) {
	posts, err = pi.PostRepository.FindByUserID(userID)
	return
}

// PostsByUserID this func is from controller to repository.
func (pi *PostInteractor) PostsByUserIDs(users domain.ResUsers, postType string) (posts domain.Posts, err error) {
	var userIDs []string
	for _, u := range users {
		userIDs = append(userIDs, u.UserID)
	}
	posts, err = pi.PostRepository.FindsConditions(
		"`type` = ? AND `user_id` IN (?)",
		postType,
		userIDs,
	)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	return
}

// PostsByTypeUserID this func is from controller to repository.
func (pi *PostInteractor) PostsByTypeUserID(userID string) (short domain.Posts, long domain.Posts, err error) {
	s := domain.Post{
		UserID: userID,
		Type:   domain.TYPE_SHORT,
	}
	l := domain.Post{
		UserID: userID,
		Type:   domain.TYPE_LONG,
	}
	short, err = pi.PostRepository.FindsConditions(s)
	long, err = pi.PostRepository.FindsConditions(l)
	return
}

// NumUserPost this func is a number of user's post.
func (pi *PostInteractor) NumUserPost(userID string) (num int, err error) {
	post := domain.Post{
		UserID: userID,
	}
	num, err = pi.PostRepository.CountConditions(post)
	return
}

// Add this func is from controller to repository.
func (pi *PostInteractor) Add(p domain.Post) (post domain.Post, err error) {
	post, err = pi.PostRepository.Store(p)
	return
}

// Update this func is from controller to repository.
func (pi *PostInteractor) Update(p domain.Post) (post domain.Post, err error) {
	post, err = pi.PostRepository.Update(p)
	return
}

// UpdateValue this func is from controller to repository.
func (pi *PostInteractor) UpdateValue(id int, column string, data string) (post domain.Post, err error) {
	post, err = pi.PostRepository.FindByID(id)
	post, err = pi.PostRepository.UpdateValue(post, column, data)
	return
}

// DeleteByID this func is from controller to repository.
func (pi *PostInteractor) DeleteByID(id int) (err error) {
	post, err := pi.PostRepository.FindByID(id)
	err = pi.PostRepository.DeleteByID(post)
	return
}
