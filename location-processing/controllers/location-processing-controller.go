package controllers

import (
	"net/http"

	service "location-processing/services"

	pb "example.com/protobuff"
	"google.golang.org/protobuf/proto"

	"github.com/gin-gonic/gin"
)

type LocationProcessingController struct {
	LocationProcessingService service.LocationProcessingInterface
}

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// createUser godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user   body      models.LocationProcessing  true  "LocationProcessing with new coordiates"
// @Success      200  {object}  models.LocationProcessing
// @Failure			 400	{string}	string	"ok"
// @Failure			 404	{string}	string	"ok"
// @Failure			 500	{string}	string	"ok"
// @Router       /user/location/ [POST]
func (u LocationProcessingController) GetDistanceTravelled(g *gin.Context) {
	user := g.Param("userID")
	startDate := g.Query("start_date")
	endDate := g.Query("end_date")

	distanceTravelled := u.LocationProcessingService.GetDistanceTravelled(user, startDate, endDate)
	g.JSON(http.StatusOK, distanceTravelled)
}

// createUser godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user   body      models.LocationProcessing  true  "LocationProcessing with new coordiates"
// @Success      200  {object}  models.LocationProcessing
// @Failure			 400	{string}	string	"ok"
// @Failure			 404	{string}	string	"ok"
// @Failure			 500	{string}	string	"ok"
// @Router       /user/location/ [POST]
func (u LocationProcessingController) UserLocationEntry(g *gin.Context) {
	// Get the raw body
	buf := make([]byte, g.Request.ContentLength)
	g.Request.Body.Read(buf)

	// Unmarshal the protobuf
	locationRequest := &pb.LocationRequest{}
	err := proto.Unmarshal(buf, locationRequest)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Could not unmarshal protobuf"})
		return
	}

	u.LocationProcessingService.StoreUserEntry(locationRequest)
}
