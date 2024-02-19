package dao

type Note struct {
	ID             int      `gorm:"column:id; primary_key; not null" json:"id"`
	UserID         int      `gorm:"column:user_id" json:"user_id"`
	DiaryID        int      `gorm:"column:diary_id" json:"diary_id"`
	PhotoCloudsUrl []string `gorm:"column:photo_clouds_url; type:text[]" json:"photo_clouds_url"`
	NatureEvents   []string `gorm:"column:nature_events; type:text[]" json:"nature_events"`
	Temperature    int      `gorm:"column:temperature" json:"temperature"`
	Supplement     string   `gorm:"column:supplement" json:"supplement"`
	LocationX      float32  `gorm:"column:location_x" json:"location_x"`
	LocationY      float32  `gorm:"column:location_y" json:"location_y"`
	AddressText    string   `gorm:"column:address_text" json:"address_text"`
	BaseModel
}
