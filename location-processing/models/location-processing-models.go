package models

type LocationProcessing struct {
	Username string  `json:"username"`
	X        float32 `json:"x" gorm:"type:float"`
	Y        float32 `json:"y" gorm:"type:float"`
	Date     string  `json:"date"`
}

type DistanceTravelledResponse struct {
	Username          string `json:"username"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	DistanceTravelled string `json:"distance_travelled"`
}
