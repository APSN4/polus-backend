package service

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"os"
	"polus-backend/app/constant"
	"polus-backend/app/domain/dao"
	"polus-backend/app/pkg"
	"polus-backend/app/repository"
	"reflect"
	"strconv"
	"syscall"
	"time"
)

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	UpdateComponentUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
	UploadPhotoUser(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
	diaryService   DiaryService
}

func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	data.Name = request.Name
	data.Surname = request.Surname
	data.UserStatusText = request.UserStatusText
	data.PhotoUrl = request.PhotoUrl
	data.DiaryID = request.DiaryID
	data.Email = request.Email
	data.Password = request.Password
	data.Status = request.Status
	data.RoleID = request.RoleID
	data.UpdatedAt = time.Now()
	u.userRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) UpdateComponentUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to update component current user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	fields := reflect.VisibleFields(reflect.TypeOf(request))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	dataDiary, err := u.diaryService.AddDiaryData(c)

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)
	request.DiaryID = dataDiary.ID
	request.CreatedAt = time.Now()

	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) GetAllUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func (u UserServiceImpl) UploadPhotoUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start upload photo by user id")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Error("Happened Error when try get file. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	err = os.Mkdir("photo_users", os.ModeDir)
	if err != nil && errors.Is(err, syscall.ERROR_ALREADY_EXISTS) {
		log.Info("Directory 'photo_users' has already exists")
	}
	if err != nil && !errors.Is(err, syscall.ERROR_ALREADY_EXISTS) {
		log.Error("Happened Error when try create folder. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	userID, _ := strconv.Atoi(c.Param("userID"))
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error("Happened Error when try create buf. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}
	err = os.WriteFile("photo_users/"+strconv.Itoa(userID)+"_"+header.Filename, buf.Bytes(), 0660)
	if err != nil {
		log.Error("Happened Error when try save file. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	type data struct {
		Filename string `json:"filename"`
	}

	dataOutput := data{Filename: strconv.Itoa(userID) + "_" + header.Filename}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, dataOutput))
}

func UserServiceInit(userRepository repository.UserRepository, diaryService DiaryService) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
		diaryService:   diaryService,
	}
}
