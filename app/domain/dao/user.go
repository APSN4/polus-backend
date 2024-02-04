package dao

type User struct {
	ID         int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Surname    string `gorm:"column:surname" json:"surname"`
	UserStatus string `gorm:"column:user_status" json:"user_status"`
	Photo      string `gorm:"column:photo_url" json:"photo_url"`
	DiaryID    int    `gorm:"column:diary_id" json:"diary_id"`
	TeachersID []int  `gorm:"column:teachers_id" json:"teachers_id"`
	ReportsID  []int  `gorm:"column:reports" json:"reports_id"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
	Status     int    `gorm:"column:status" json:"status"`
	RoleID     int    `gorm:"column:role_id;not null" json:"role_id"`
	Role       Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
