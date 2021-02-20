package database

import (
	"github.com/maaaaakoto35/PostUpAPI/domain"
)

// UserRepository this struct has SQLHandler.
type UserRepository struct {
	SQLHandler
}

// FindByID this func is finding user by id.
func (userRepository *UserRepository) FindByID(id int) (user domain.User, err error) {
	if err = userRepository.Find(&user, id).Error; err != nil {
		return
	}
	return
}

// Store this func is storing user.
func (userRepository *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = userRepository.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// Update this func is updating user.
func (userRepository *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = userRepository.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// DeleteByID this func is deletinguser by id.
func (userRepository *UserRepository) DeleteByID(user domain.User) (err error) {
	if err = userRepository.Delete(&user).Error; err != nil {
		return
	}
	return
}

// FindAll this func is finding all users.
func (userRepository *UserRepository) FindAll() (users domain.Users, err error) {
	if err = userRepository.Find(&users).Error; err != nil {
		return
	}
	return
}
