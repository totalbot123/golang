package models

type Users struct {
	Username string  `json:"username" gorm:"primary_key,column:username"`
	X        float32 `json:"x" gorm:"type:float"`
	Y        float32 `json:"y" gorm:"type:float"`
}
