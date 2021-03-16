package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// FollowRepository this interface is connecting FollowRepository.
type FollowRepository interface {
	FindConditions(...interface{}) (domain.Follow, error)
	Store(domain.Follow) (domain.Follow, error)
	Update(domain.Follow) (domain.Follow, error)
	UpdateValue(domain.Follow, string, string) (domain.Follow, error)
	DeleteByID(domain.Follow) error
}
