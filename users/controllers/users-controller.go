package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	models "users/models"
	service "users/services"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService service.UsersInterface
}

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// createUsers godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user   body      models.Users  true  "Users with new coordiates"
// @Success      200  {object}  models.Users
// @Failure			 400	{string}	string	"ok"
// @Failure			 404	{string}	string	"ok"
// @Failure			 500	{string}	string	"ok"
// @Router       /user/location/ [POST]
func (u UsersController) CreateUsers(g *gin.Context) {
	var user models.Users

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = u.UsersService.CreateUsers(user.Username, user.X, user.Y)
	g.JSON(http.StatusOK, user)
}

// updateUsers godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        user   body      models.Users  true  "Users with new coordiates"
// @Success      200  {object}  models.Users
// @Failure			 400	{string}	string	"ok"
// @Failure			 404	{string}	string	"ok"
// @Failure			 500	{string}	string	"ok"
// @Router       /user/location/ [PATCH]
func (u UsersController) UpdateUsers(g *gin.Context) {
	var user models.Users

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.UsersService.UpdateUsers(user.Username, user.X, user.Y)
	g.JSON(http.StatusOK, user)
}

// getUserss godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        x   path      number  true  "X coordinate"
// @Param        y   path      number  true  "Y coordinate"
// @Success      200  {array}  models.Users
// @Failure			 400	{string}	string	"ok"
// @Failure			 404	{string}	string	"ok"
// @Failure			 500	{string}	string	"ok"
// @Router       /user/ [GET]
func (u UsersController) GetUsers(g *gin.Context) {
	x, err := strconv.ParseFloat(g.Query("x"), 32)
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.ParseFloat(g.Query("y"), 32)
	if err != nil {
		fmt.Println(err)
	}

	radius, err := strconv.ParseFloat(g.Query("radius"), 32)
	if err != nil {
		fmt.Println(err)
	}

	users := u.UsersService.GetUsers(float32(x), float32(y), float32(radius))
	g.JSON(http.StatusOK, users)
}
