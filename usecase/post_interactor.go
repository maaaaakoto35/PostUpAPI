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
