package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// PostRepository this interface is connecting PostRepository.
type PostRepository interface {
	FindByID(id int) (domain.Post, error)
	FindByUserID(string) (domain.Post, error)
	FindConditions(...interface{}) (domain.Post, error)
	CountConditions(...interface{}) (int, error)
	Store(domain.Post) (domain.Post, error)
	Update(domain.Post) (domain.Post, error)
	UpdateValue(domain.Post, string, string) (domain.Post, error)
	DeleteByID(domain.Post) error
}
