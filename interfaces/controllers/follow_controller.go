package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/domain"
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

// FollowedGetImpl this func is inisializing FollowController.
func (controller *FollowController) FollowedGetImpl(c Context) (follower domain.ResUsers, err error) {
	userID := jwtUserID(c)
	follower, err = controller.Interactor.FollowedUserID(userID)
	return
}

// FollowingGetImpl this func is inisializing FollowController.
func (controller *FollowController) FollowingGetImpl(c Context) (follows domain.ResUsers, err error) {
	userID := jwtUserID(c)
	follows, err = controller.Interactor.FollowingUserID(userID)
	return
}
