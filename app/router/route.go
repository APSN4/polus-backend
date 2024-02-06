package router

import (
	"github.com/gin-gonic/gin"
	"polus-backend/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		user := api.Group("/user")
		user.GET("", init.UserCtrl.GetAllUserData)
		user.POST("", init.UserCtrl.AddUserData)
		user.GET("/:userID", init.UserCtrl.GetUserById)
		user.PUT("/:userID", init.UserCtrl.UpdateUserData)
		user.DELETE("/:userID", init.UserCtrl.DeleteUser)

		diary := api.Group("/diary")
		diary.GET("", init.DiaryCtrl.GetAllDiaryData)
		diary.POST("", init.DiaryCtrl.AddDiaryData)
		diary.GET("/:diaryID", init.UserCtrl.GetUserById)
		diary.PUT("/:diaryID", init.UserCtrl.UpdateUserData)
		diary.DELETE("/:diaryID", init.UserCtrl.DeleteUser)
	}

	return router
}
