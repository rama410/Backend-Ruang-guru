package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	// TODO: answer here
	var userLogin model.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid decode json"})
		return
	}

	user := model.User{
		Email:    userLogin.Email,
		Password: userLogin.Password,
	}

	token, err := u.userService.Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "login fail"})
		return
	}

	c.SetCookie("session_token", *token, 1000, "/", "localhost", false, true)

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "login success"})
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	// TODO: answer here
	tc, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "task category not found"})
		return
	}else {
		c.JSON(http.StatusOK, tc)
	}
}
