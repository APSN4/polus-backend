package dao

type User struct {
	ID              int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name            string `gorm:"column:name" json:"name"`
	Surname         string `gorm:"column:surname" json:"surname"`
	UserStatusText  string `gorm:"column:user_status_text" json:"user_status_text"`
	PhotoUrl        string `gorm:"column:photo_url" json:"photo_url"`
	DiaryID         int    `gorm:"column:diary_id" json:"diary_id"`
	Diary           *Diary `gorm:"foreignKey:DiaryID;references:ID" json:"diary"`
	DiaryStudentsID []int  `gorm:"column:diary_students_id; type:int[]" json:"diary_students_id"`
	TeachersID      []int  `gorm:"column:teachers_id; type:int[]" json:"teachers_id"`
	NotesID         []int  `gorm:"column:notes_id; type:int[]" json:"notes_id"`
	Notes           []Note `gorm:"foreignKey:UserID" json:"notes"`
	Email           string `gorm:"column:email" json:"email"`
	Password        string `gorm:"column:password" json:"password"`
	Status          int    `gorm:"column:status" json:"status"`
	RoleID          int    `gorm:"column:role_id;not null" json:"role_id"`
	Role            Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
