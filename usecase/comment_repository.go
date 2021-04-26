package usecase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// CommentRepository this interface is connecting CommentRepository.
type CommentRepository interface {
	FindByID(int) (domain.Comment, error)
	FindByUserID(string) (domain.Comment, error)
	Store(domain.Comment) (domain.Comment, error)
	Update(domain.Comment) (domain.Comment, error)
	DeleteByID(domain.Comment) (domain.Comment, error)
}
