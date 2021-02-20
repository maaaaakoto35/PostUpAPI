package controllers

import (
	"net/http"
	"strconv"

	"github.com/maaaaakoto35/PostUpAPI/domain"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// UserController this struct is recieving Interactor interface.
type UserController struct {
	Interactor usecase.UserInteractor
}

// NewUserController this func is initializing UserController.
func NewUserController(sqlHandler database.SQLHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// CreateUser this func is creating user.
func (controller *UserController) CreateUser(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	user, err := controller.Interactor.Add(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}

// GetUsers this func is getting users.
func (controller *UserController) GetUsers(c Context) (err error) {
	users, err := controller.Interactor.Users()

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

// GetUser this func is getting a user.
func (controller *UserController) GetUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

// UpdateUser this func is updating user.
func (controller *UserController) UpdateUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	u := domain.User{ID: id}
	c.Bind(&u)

	user, err := controller.Interactor.Update(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}

// DeleteUser this func is deleting user.
func (controller *UserController) DeleteUser(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := domain.User{ID: id}

	err = controller.Interactor.DeleteByID(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
