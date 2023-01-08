package services

import (
	pb "example.com/protobuff"
	models "location-processing/models"
)

type LocationProcessingInterface interface {
	GetDistanceTravelled(user, startDate, endDate string) models.DistanceTravelledResponse
	StoreUserEntry(locationRequest *pb.LocationRequest)
}
