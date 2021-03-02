package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/domain"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// PostController this struct is recieving Interactor interface.
type PostController struct {
	Interactor usecase.PostInteractor
}

// NewPostController this func is initializing PostController.
func NewPostController(sqlHandler database.SQLHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// CreatePost this func is creating post.
func (controller *PostController) CreatePost(c Context) (err error) {
	p := domain.Post{}
	c.Bind(&p)
	user, err := controller.Interactor.Add(p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}
