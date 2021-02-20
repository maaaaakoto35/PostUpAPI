package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// UserInteractor this struct has UserRepository.
type UserInteractor struct {
	UserRepository UserRepository
}

// UserByID this func is from controller to repository.
func (interactor *UserInteractor) UserByID(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

// Users this func is from controller to repository.
func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

// Add this func is from controller to repository.
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

// Update this func is from controller to repository.
func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

// DeleteByID this func is from controller to repository.
func (interactor *UserInteractor) DeleteByID(user domain.User) (err error) {
	err = interactor.UserRepository.DeleteByID(user)
	return
}
