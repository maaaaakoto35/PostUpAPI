package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// PostInteractor this struct has PostRepository.
type PostInteractor struct {
	PostRepository PostRepository
}

// PostByID this func is from controller to repository.
func (pi *PostInteractor) PostByID(id int) (post domain.Post, err error) {
	post, err = pi.PostRepository.FindByID(id)
	return
}

// PostByUserID this func is from controller to repository.
func (pi *PostInteractor) PostByUserID(userID string) (post domain.Post, err error) {
	post, err = pi.PostRepository.FindByUserID(userID)
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
