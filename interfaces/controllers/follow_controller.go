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

// FFnumImpl this func is getting ff nums for impl.
func (controller *FollowController) FfNumImpl(c Context, userID string) (following int, followed int, err error) {
	if userID == "" {
		userID = jwtUserID(c)
	}
	following, err = controller.Interactor.FollowingNum(userID)
	followed, err = controller.Interactor.FollowedNum(userID)
	return
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

//
func (controller *FollowController) Follow(c Context) (err error) {
	f := domain.Follow{}
	c.Bind(&f)
	follow, err := controller.Interactor.Add(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, follow)
	return
}

//
func (controller *FollowController) UnFollow(c Context) (err error) {
	f := domain.Follow{}
	c.Bind(&f)
	err = controller.Interactor.Delete(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
