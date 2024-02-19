package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"polus-backend/app/domain/dao"
)

type DiaryRepository interface {
	FindAllDiary() ([]dao.Diary, error)
	FindDiaryById(id int) (dao.Diary, error)
	Save(diary *dao.Diary) (dao.Diary, error)
	DeleteDiaryById(id int) error
}

type DiaryRepositoryImpl struct {
	db *gorm.DB
}

func (d DiaryRepositoryImpl) FindAllDiary() ([]dao.Diary, error) {
	var diaries []dao.Diary

	var err = d.db.Preload("User").Find(&diaries).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return diaries, nil
}

func (d DiaryRepositoryImpl) FindDiaryById(id int) (dao.Diary, error) {
	diary := dao.Diary{
		ID: id,
	}
	err := d.db.Preload("User").First(&diary).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.Diary{}, err
	}
	return diary, nil
}

func (d DiaryRepositoryImpl) Save(diary *dao.Diary) (dao.Diary, error) {
	var err = d.db.Save(diary).Error
	if err != nil {
		log.Error("Got an error when save diary. Error: ", err)
		return dao.Diary{}, err
	}
	return *diary, nil
}

func (d DiaryRepositoryImpl) DeleteDiaryById(id int) error {
	err := d.db.Delete(&dao.Diary{}, id).Error
	if err != nil {
		log.Error("Got an error when delete diary. Error: ", err)
		return err
	}
	return nil
}

func DiaryRepositoryInit(db *gorm.DB) *DiaryRepositoryImpl {
	db.AutoMigrate(&dao.Diary{})
	return &DiaryRepositoryImpl{
		db: db,
	}
}
