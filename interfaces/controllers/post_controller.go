package controllers

import (
	"net/http"

	"github.com/maaaaakoto35/PostUpAPI/domain"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/database"
	"github.com/maaaaakoto35/PostUpAPI/interfaces/storage"
	"github.com/maaaaakoto35/PostUpAPI/usecase"
)

// PostController this struct is recieving Interactor interface.
type PostController struct {
	Interactor usecase.PostInteractor
}

// StorageController this struct is recieving interface.
type StorageController struct {
	StorageController storage.StorageHandler
}

// NewPostController this func is initializing PostController.
func NewPostController(sqlHandler database.SQLHandler, storageHandler storage.StorageHandler) (db *PostController, storage *StorageController) {
	db = &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SQLHandler: sqlHandler,
			},
		},
	}

	storage = &StorageController{
		StorageController: storageHandler,
	}

	return
}

// GetFederation this func is response token.
func (controller *StorageController) GetFederation(c Context) (err error) {
	userID := jwtUserID(c)
	federationToken, err := controller.StorageController.GetFederationToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, federationToken)
	return
}

// GetPresignedURL this func is getting pre-sign url.
func (controller *StorageController) GetPresignedURL(c Context) (err error) {
	userID := jwtUserID(c)
	num := c.Get("num").(int)
	url, err := controller.StorageController.GetPresignedURL(userID, num)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status string `json:"status"`
		URL    string `json:"url"`
	}{
		Status: "success",
		URL:    url,
	})
	return
}

// CreatePost this func is creating post.
func (controller *PostController) CreatePost(c Context) (err error) {
	p := domain.Post{}
	c.Bind(&p)
	post, err := controller.Interactor.Add(p)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, post)
	return
}

// GetUserPost this func is getting user posts.
func (controller *PostController) GetUserPost(c Context) (err error) {
	userID := c.Param("user_id")
	posts, err := controller.Interactor.PostsByUserID(userID)

	// comment
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, posts)
	return
}

// GetUserPostNumImpl this func is getting only num.
func (controller *PostController) GetUserPostNumImpl(c Context) (num int, err error) {
	userID := jwtUserID(c)
	num, err = controller.Interactor.NumUserPost(userID)
	return
}

// GetUserPostNumImpl this func is getting only num.
func (controller *PostController) GetPostsImpl(c Context, userID string) (short domain.Posts, long domain.Posts, err error) {
	if userID == "" {
		userID = jwtUserID(c)
	}
	short, long, err = controller.Interactor.PostsByTypeUserID(userID)
	return
}

// GetShort this func is getting short for home.
func (controller *PostController) GetPostsFollowing(c Context, following domain.ResUsers, postType string) (err error) {
	posts, err := controller.Interactor.PostsByUsers(following, postType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, posts)
	return
}

// GetPostNum this func is getting post upping num.
func (controller *PostController) GetPostNum(c Context) (err error) {
	userID := c.Param("user_id")
	num, err := controller.Interactor.NumUserPost(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status  string `json:"status"`
		PostNum int    `json:"post_num"`
	}{
		Status:  "success",
		PostNum: num,
	})
	return
}

// WatchPost this func is watching post.
func (controller *PostController) WatchPost(c Context) (err error) {
	// userID := jwtUserID(c)
	post := domain.Post{}
	c.Bind(&post)
	post.Watch++
	_, err = controller.Interactor.Update(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})
	return
}

// GoodPost this func is doing good post.
func (controller *PostController) GoodPost(c Context) (err error) {
	// userID := jwtUserID(c)
	post := domain.Post{}
	c.Bind(&post)
	post.Good++
	_, err = controller.Interactor.Update(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusAccepted, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})
	return
}
