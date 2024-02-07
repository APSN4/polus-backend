package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"polus-backend/app/domain/dao"
)

type NoteRepository interface {
	FindAllNote() ([]dao.Note, error)
	FindNoteById(id int) (dao.Note, error)
	Save(note *dao.Note) (dao.Note, error)
	DeleteNoteById(id int) error
}

type NoteRepositoryImpl struct {
	db *gorm.DB
}

func (d NoteRepositoryImpl) FindAllNote() ([]dao.Note, error) {
	var note []dao.Note

	var err = d.db.Find(&note).Error // Preload("Note")
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return note, nil
}

func (d NoteRepositoryImpl) FindNoteById(id int) (dao.Note, error) {
	note := dao.Note{
		ID: id,
	}
	err := d.db.Preload("User").First(&note).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.Note{}, err
	}
	return note, nil
}

func (d NoteRepositoryImpl) Save(note *dao.Note) (dao.Note, error) {
	var err = d.db.Save(note).Error
	if err != nil {
		log.Error("Got an error when save note. Error: ", err)
		return dao.Note{}, err
	}
	return *note, nil
}

func (d NoteRepositoryImpl) DeleteNoteById(id int) error {
	err := d.db.Delete(&dao.Note{}, id).Error
	if err != nil {
		log.Error("Got an error when delete note. Error: ", err)
		return err
	}
	return nil
}

func NoteRepositoryInit(db *gorm.DB) *NoteRepositoryImpl {
	db.AutoMigrate(&dao.Note{})
	return &NoteRepositoryImpl{
		db: db,
	}
}
