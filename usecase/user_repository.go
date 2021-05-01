package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// UserRepository this interface is connecting UserRepository.
type UserRepository interface {
	FindByID(id int) (domain.User, error)
	FindByUserID(string) (domain.User, error)
	FindConditions(...interface{}) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	UpdateValue(domain.User, string, string) (domain.User, error)
	DeleteByID(domain.User) error
	FindAll() (domain.Users, error)
}
