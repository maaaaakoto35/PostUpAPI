package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// CommentInteractor this struct is has CommentRepository.
type CommentInteractor struct {
	CommentRepository CommentRepository
}

// CommentByID this func is from controller to repository.
func (ci *CommentInteractor) CommentByID(id int) (comment domain.Comment, err error) {
	comment, err = ci.CommentRepository.FindByID(id)
	return
}

// Add this func is from controller to repository.
func (ci *CommentInteractor) Add(c domain.Post) (comment domain.Comment, err error) {
	comment, err = ci.CommentRepository.Store(c)
	return
}

// Update this func is from controller to repository.
func (ci *CommentInteractor) Update(c domain.Comment) (comment domain.Comment, err error) {
	comment, err = ci.CommentRepository.Update(c)
	return
}

// UpdateValue this func is from controller to repository.
func (ci *CommentInteractor) UpdateValue(id int, column string, data string) (comment domain.Comment, err error) {
	comment, err = ci.CommentRepository.FindByID(id)
	comment, err = ci.CommentRepository.UpdateValue(comment, column, data)
	return
}

// DeleteByID this func is from controller to repository.
func (ci *CommentInteractor) DeleteByID(id int) (err error) {
	comment, err := ci.CommentRepository.FindByID(id)
	err = ci.CommentRepository.DeleteByID(comment)
	return
}
