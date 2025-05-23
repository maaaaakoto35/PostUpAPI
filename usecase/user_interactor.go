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

// CanLogin this func is existing user_id and password
func (interactor *UserInteractor) CanLogin(userID, password string) (bool, error) {
	_, err := interactor.UserRepository.FindConditions(
		domain.User{
			UserID: userID,
			Pass:   password,
		},
	)
	if err != nil {
		return false, err
	}
	return true, err
}

// ResUserByUserID this func is from controller to repository
func (interactor *UserInteractor) ResUserByUserID(userID string) (resUser domain.ResUser, err error) {
	user, err := interactor.UserRepository.FindByUserID(userID)
	resUser = domain.BindUser(user)
	return
}

// ResUsersByResUsers this func is to get full resusers.
func (interactor *UserInteractor) ResUsersByResUsers(res domain.ResUsers) (resUsers domain.ResUsers, err error) {
	for _, r := range res {
		user, e := interactor.UserRepository.FindByUserID(r.UserID)
		resUser := domain.BindUser(user)

		resUsers = append(resUsers, resUser)
		err = e
	}

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

// UpdateValue this func is from controller to repository.
func (interactor *UserInteractor) UpdateValue(userID string, column string, data string) (resUser domain.ResUser, err error) {
	user, err := interactor.UserRepository.FindByUserID(userID)
	user, err = interactor.UserRepository.UpdateValue(user, column, data)
	resUser = domain.BindUser(user)
	return
}

// DeleteByID this func is from controller to repository.
func (interactor *UserInteractor) DeleteByID(userID string) (err error) {
	user, err := interactor.UserRepository.FindByUserID(userID)
	err = interactor.UserRepository.DeleteByID(user)
	return
}
