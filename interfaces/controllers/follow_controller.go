package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// FollowController this struct is recieving Interactor interface.
type FollowController struct {
	Interactor usecase.FollowInteractor
}

// NewFollowController this func is initializing FollowController.
func NewFollowController(sqlHandler database.SQLHandler) *FollowController {
	return &FollowController{
		Interactor: usecase.FollowInteractor{
			FollowRepository: &database.FollowRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// Followed this func is initializing FollowController.
func (controller *FollowController) Followed(c Context) (err error) {
	userID := jwtUserID(c)
	follow, err := controller.Interactor.FollowedUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, follow)
	return
}

// Following this func is initializing FollowController.
func (controller *FollowController) Following(c Context) (err error) {
	userID := jwtUserID(c)
	follow, err := controller.Interactor.FollowingUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, follow)
	return
}
