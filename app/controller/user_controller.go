package controller

import (
	"github.com/gin-gonic/gin"
	"polus-backend/app/service"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	UpdateComponentUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	UploadPhotoUser(c *gin.Context)
}

type UserControllerImpl struct {
	svc service.UserService
}

func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) UpdateComponentUserData(c *gin.Context) {
	u.svc.UpdateComponentUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func (u UserControllerImpl) UploadPhotoUser(c *gin.Context) {
	u.svc.UploadPhotoUser(c)
}

func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
