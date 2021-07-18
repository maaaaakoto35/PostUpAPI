package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/domain"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// UserController this struct is recieving Interactor interface.
type UserController struct {
	Interactor usecase.UserInteractor
}

// UpdateValue this struct is recieving posting data.
type UpdateValue struct {
	Column string `json:"column"`
	Data   string `json:"data"`
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

// LogIn this func is logging in.
func (controller *UserController) LogIn(c Context) (err error) {
	u := domain.User{}
	c.Bind(&u)

	// varidattion
	if u.UserID == "" || u.Pass == "" {
		c.JSON(http.StatusInternalServerError, "dose not match args")
		return
	}

	result, err := controller.Interactor.CanLogin(u.UserID, u.Pass)
	if err != nil || result != true {
		c.JSON(http.StatusNonAuthoritativeInfo, NewError(err))
		return
	}

	// set custom claims
	token, err := setJwt(u.UserID, u.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}{
		Status: "success",
		Token:  token,
	})
	return
}

// GetInfo this func is getting a user.
func (controller *UserController) GetInfo(c Context, userID string) (err error) {
	if userID == "" {
		userID = jwtUserID(c)
	}
	user, err := controller.Interactor.ResUserByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	follower := c.Get("follower").(int)
	following := c.Get("following").(int)
	short, ok := c.Get("short").(*domain.Posts)
	if !ok {
		short = nil
	}
	long, ok := c.Get("long").(*domain.Posts)
	if !ok {
		long = nil
	}
	param := domain.BindParam{
		ResUser:   user,
		Following: following,
		Follower:  follower,
		Short:     short,
		Long:      long,
	}
	resUser := domain.Bind(param)
	c.JSON(http.StatusOK, resUser)
	return
}

// GetUsers this func is getting users.
func (controller *UserController) GetUsers(c Context) (err error) {
	users, err := controller.Interactor.ResUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

// GetUser this func is getting a user.
func (controller *UserController) GetUser(c Context) (err error) {
	userID := c.Param("user_id")
	user, err := controller.Interactor.ResUserByUserID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

// GetUserImpl this func is getting a user.
func (controller *UserController) GetUserImpl(userID string) (user domain.ResUser, err error) {
	user, err = controller.Interactor.ResUserByUserID(userID)
	return
}

// UpdateUser this func is updating user.
func (controller *UserController) UpdateUser(c Context) (err error) {
	userID := jwtUserID(c)
	updateValue := new(UpdateValue)
	c.Bind(updateValue)

	user, err := controller.Interactor.UpdateValue(userID, updateValue.Column, updateValue.Data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, user)
	return
}

// DeleteUser this func is deleting user.
func (controller *UserController) DeleteUser(c Context) (err error) {
	userID := jwtUserID(c)

	err = controller.Interactor.DeleteByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}

// ResFollows this func is responce follows.
func (controller *UserController) ResFollows(c Context) (err error) {
	follows := c.Get("follows").(domain.ResUsers)
	resUsers, err := controller.Interactor.ResUsersByResUsers(follows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, resUsers)
	return
}
