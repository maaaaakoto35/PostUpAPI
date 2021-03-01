package database

import (
	"github.com/maaaaakoto35/PostUpAPI/domain"
)

// UserRepository this struct has SQLHandler.
type UserRepository struct {
	SQLHandler
}

// FindByID this func is finding user by id.
func (ur *UserRepository) FindByID(id int) (user domain.User, err error) {
	if err = ur.Find(&user, id).Error; err != nil {
		return
	}
	return
}

// FindByUserID this func is finding user by user_id.
func (ur *UserRepository) FindByUserID(userID string) (user domain.User, err error) {
	u := domain.User{
		UserID: userID,
	}
	if err = ur.Find(&user, u).Error; err != nil {
		return
	}
	return
}

// FindConditions this func is finding user by some conditions.
func (ur *UserRepository) FindConditions(where ...interface{}) (user domain.User, err error) {
	if err = ur.Find(&user, where...).Error; err != nil {
		return
	}
	return
}

// Store this func is storing user.
func (ur *UserRepository) Store(u domain.User) (user domain.User, err error) {
	if err = ur.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// Update this func is updating user.
func (ur *UserRepository) Update(u domain.User) (user domain.User, err error) {
	if err = ur.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// UpdateValue this func is updating some columns in user.
func (ur *UserRepository) UpdateValue(u domain.User, set string, value string) (user domain.User, err error) {
	if err = ur.SaveValue(&u, set, value).Error; err != nil {
		return
	}
	user = u
	return
}

// DeleteByID this func is deletinguser by id.
func (ur *UserRepository) DeleteByID(user domain.User) (err error) {
	if err = ur.Delete(&user).Error; err != nil {
		return
	}
	return
}

// FindAll this func is finding all users.
func (ur *UserRepository) FindAll() (users domain.Users, err error) {
	if err = ur.Find(&users).Error; err != nil {
		return
	}
	return
}
