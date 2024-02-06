package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"polus-backend/app/constant"
	"polus-backend/app/domain/dao"
	"polus-backend/app/pkg"
	"polus-backend/app/repository"
	"strconv"
)

type DiaryService interface {
	GetAllDiary(c *gin.Context)
	GetDiaryById(c *gin.Context)
	AddDiaryData(c *gin.Context)
	UpdateDiaryData(c *gin.Context)
	DeleteDiary(c *gin.Context)
}

type DiaryServiceImpl struct {
	diaryRepository repository.DiaryRepository
}

func (u DiaryServiceImpl) UpdateDiaryData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update diary data by id")
	diaryID, _ := strconv.Atoi(c.Param("diaryID"))

	var request dao.Diary
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.diaryRepository.FindDiaryById(diaryID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	u.diaryRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u DiaryServiceImpl) GetDiaryById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get diary by id")
	diaryID, _ := strconv.Atoi(c.Param("diaryID"))

	data, err := u.diaryRepository.FindDiaryById(diaryID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u DiaryServiceImpl) AddDiaryData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data diary")
	var request dao.Diary
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.diaryRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u DiaryServiceImpl) GetAllDiary(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data Diary")

	data, err := u.diaryRepository.FindAllDiary()
	if err != nil {
		log.Error("Happened Error when find all diary data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u DiaryServiceImpl) DeleteDiary(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data Diary by id")
	DiaryID, _ := strconv.Atoi(c.Param("diaryID"))

	err := u.diaryRepository.DeleteDiaryById(DiaryID)
	if err != nil {
		log.Error("Happened Error when try delete data Diary from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func DiaryServiceInit(diaryRepository repository.DiaryRepository) *DiaryServiceImpl {
	return &DiaryServiceImpl{
		diaryRepository: diaryRepository,
	}
}
