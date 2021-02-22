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

// ResUserByID this func is from controller to repository
func (interactor *UserInteractor) ResUserByID(id int) (resUser domain.ResUser, err error) {
	user, err := interactor.UserRepository.FindByID(id)
	resUser = domain.BindUser(user)
	return
}

// ResUsers this func is from controller to repository.
func (interactor *UserInteractor) ResUsers() (resUsers domain.ResUsers, err error) {
	users, err := interactor.UserRepository.FindAll()
	resUsers = domain.BindUsers(users)
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
