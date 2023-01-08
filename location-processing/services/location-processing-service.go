package services

import (
	"fmt"
	"location-processing/dao"
	models "location-processing/models"
	"math"
	"time"

	pb "example.com/protobuff"
	// "location-processing/util"
)

type LocationProcessingServiceProvider struct{}

func NewLocationProcessingService() *LocationProcessingServiceProvider {
	return &LocationProcessingServiceProvider{}
}

func (u LocationProcessingServiceProvider) GetDistanceTravelled(user, startDate, endDate string) models.DistanceTravelledResponse {
	var databaseEntries []models.LocationProcessing

	dao.DB.Where("date > ?", startDate).Where("date < ?", endDate).Where("username = ?", user).Find(&databaseEntries)
	distanceTravelled := fmt.Sprintf("%.2f", getDistanceTravelled(databaseEntries))

	return models.DistanceTravelledResponse{
		Username:          user,
		StartDate:         startDate,
		EndDate:           endDate,
		DistanceTravelled: distanceTravelled,
	}
}

func getDistanceTravelled(entires []models.LocationProcessing) float32 {
	var distanceTravelled float32 = 0
	if len(entires) < 2 {
		return 0
	}
	for i := 1; i < len(entires); i++ {
		distanceTravelled += euclidianDistance(entires[i-1], entires[i])
	}

	return distanceTravelled
}

func euclidianDistance(p1, p2 models.LocationProcessing) float32 {
	a := float64(p1.X - p2.X)
	b := float64(p1.Y - p2.Y)
	return float32(math.Sqrt(a*a + b*b))
}

func (u LocationProcessingServiceProvider) StoreUserEntry(locationRequest *pb.LocationRequest) {
	now := time.Now()
	formattedDate := now.Format("2006-01-02T15:04:05+00:00")
	locationProcessing := models.LocationProcessing{
		Username: locationRequest.Username,
		X:        locationRequest.X,
		Y:        locationRequest.Y,
		Date:     formattedDate,
	}

	dao.DB.Model(locationProcessing).Create(locationProcessing)
}
