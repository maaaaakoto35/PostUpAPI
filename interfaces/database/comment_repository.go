package datebase

import "github.com/maaaaakoto35/PostUpAPI/domain"

// CommentRepository this struct has SQLHandler.
type CommentRepository struct {
	SQLHandler
}

// FindByID this func is finding post by id.
func (cr *CommentRepository) FindByID(id int) (comment domain.Comment, err error) {
	if err = cr.Find(&comment, id).Error; err != nil {
		return
	}
	return
}

// FindByUserID this func is finding comment by comment_id.
func (cr *CommentRepository) FindByUserID(userID string) (comments domain.Comments, err error) {
	c := domain.Comment{
		UserID: userID,
	}
	if err = cr.Find(&comments, c).Error; err != nil {
		return
	}
	return
}

// Store this func is storing comment.
func (cr *CommentRepository) Store(c domain.Comment) (comment domain.Comment, err error) {
	if err = cr.Create(&c).Error; err != nil {
		return
	}
	comment = c
	return
}

// Update this func is updating comment.
func (cr *CommentRepository) Update(c domain.Comment) (comment domain.Comment, err error) {
	if err = cr.Save(&c).Error; err != nil {
		return
	}
	comment = c
	return
}

// DeleteByID this func is deleting comment by id.
func (cr *CommentRepository) DeleteByID(comment domain.Comment) (err error) {
	if err = cr.Delete(&comment).Error; err != nil {
		return
	}
	return
}
