package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/domain"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/storage"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// CommentController this is struct is recieving Interactor interface.
type CommentController struct {
	Interactor usecase.CommentController
}

// NewPostController this func is initializing CommentController.
func NewCommentController(sqlHandler database.SQLHandler) *CommentController {
	db = &CommentController{
		Interactor: usecase.CommentInteractor{
			CommentRepository: &database.CommentRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// CreateComment this func is write a comment.
func (controller *CommentController) CreateComment(c Context) (err error) {
	cc := domain.Comment{}
	c.Bind(&cc)
	comment, err := controller.Interactor.Add(cc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, comment)
	return
}

// DeleteComment this func is delete a comment.
func (controller *CommentController) DeleteComment(c Context) (string, err error) {
	cc := domain.Commnet{}
	c.Bind(&cc)
	comment, err := controller.Interactor.Delete(cc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, comment)
	return
}

func (controller *CommentController) Edit(c Context) (bool, err error) {
	cc := domain.Commnet{}
	c.Bind(&cc)
	comment, err := controller.Interactor.Update(cc)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, comment)
	return
}

// GoodComment this func is doing good comment.
func (controller *CommentController) GoodComment(c Context) (err error) {
	comment := domain.Comment{}
	c.Bind(&comment)
	comment.Good++
	_, err = controller.Interactor.Update(comment)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status string `"json:status"`
	}{
		Status: "success",
	})
	return
}
