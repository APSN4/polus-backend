package dao

type Diary struct {
	ID      int    `gorm:"column:id; primary_key; not null" json:"id"`
	NotesID []int  `gorm:"column:notes_id; type:int[]" json:"notes_id"`
	Notes   []Note `gorm:"foreignKey:DiaryID" json:"notes"`
	BaseModel
}
