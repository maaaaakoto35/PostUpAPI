package database

import "github.com/maaaaakoto35/PostUpAPI/domain"

// PostRepository this struct has SQLHandler.
type PostRepository struct {
	SQLHandler
}

// FindByID this func is finding post by id.
func (pr *PostRepository) FindByID(id int) (post domain.Post, err error) {
	if err = pr.Find(&post, id).Error; err != nil {
		return
	}
	return
}

// FindByUserID this func is finding post by post_id.
func (pr *PostRepository) FindByUserID(userID string) (post domain.Post, err error) {
	p := domain.Post{
		UserID: userID,
	}
	if err = pr.Find(&post, p).Error; err != nil {
		return
	}
	return
}

// FindConditions this func is finding post by some conditions.
func (pr *PostRepository) FindConditions(where ...interface{}) (post domain.Post, err error) {
	if err = pr.Find(&post, where...).Error; err != nil {
		return
	}
	return
}

// Store this func is storing post.
func (pr *PostRepository) Store(p domain.Post) (post domain.Post, err error) {
	if err = pr.Create(&p).Error; err != nil {
		return
	}
	post = p
	return
}

// Update this func is updating post.
func (pr *PostRepository) Update(p domain.Post) (post domain.Post, err error) {
	if err = pr.Save(&p).Error; err != nil {
		return
	}
	post = p
	return
}

// UpdateValue this func is updating some columns in post.
func (pr *PostRepository) UpdateValue(p domain.Post, set string, value string) (post domain.Post, err error) {
	if err = pr.SaveValue(&p, set, value).Error; err != nil {
		return
	}
	post = p
	return
}

// DeleteByID this func is deletingpost by id.
func (pr *PostRepository) DeleteByID(post domain.Post) (err error) {
	if err = pr.Delete(&post).Error; err != nil {
		return
	}
	return
}
