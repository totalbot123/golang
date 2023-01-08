package services

import (
	"bytes"
	"net/http"
	"users/dao"
	models "users/models"

	pb "example.com/protobuff"
	"google.golang.org/protobuf/proto"

	"users/util"
)

type UsersServiceProvider struct{}

func NewUsersService() *UsersServiceProvider {
	return &UsersServiceProvider{}
}

func (u UsersServiceProvider) GetUsers(x, y, radius float32) []models.Users {
	var users []models.Users
	util.Logger.Debug("")
	dao.DB.Where("sqrt(pow((x - ?), 2) +  pow((y - ?), 2)) <= ?", x, y, radius).Find(&users)

	return users
}

func (u UsersServiceProvider) CreateUsers(userName string, x, y float32) models.Users {
	user := models.Users{
		Username: userName,
		X:        x,
		Y:        y,
	}

	dao.DB.Create(&user)
	u.updateLocationHistory(user)
	return user
}

func (u UsersServiceProvider) UpdateUsers(userName string, x, y float32) models.Users {
	user := models.Users{
		Username: userName,
		X:        x,
		Y:        y,
	}

	dao.DB.Model(&models.Users{}).Where("username = ?", userName).Updates(user)
	u.updateLocationHistory(user)
	return user
}

func (u UsersServiceProvider) updateLocationHistory(user models.Users) error {
	locationRequest := pb.LocationRequest{
		Username: user.Username,
		X:        user.X,
		Y:        user.Y,
	}

	rawData, err := proto.Marshal(&locationRequest)
	if err != nil {
		return err
	}

	data := bytes.NewReader(rawData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8002/api/v1/history", data)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/protobuf")

	// Send the request
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
