package controller

import (
	"net/http"
	"problem2/config"
	"problem2/model"

	"github.com/labstack/echo"
)

// ------------------controller

// get all users
func GetUsersController(c echo.Context)error  {
	var users []model.User

	if err := config.DB.Find(&users).Error;err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,map[string]interface{}{
		"messages" : "success get all users",
		"users": users,
	})
}


// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	user := model.User{}

	config.DB.Where("Id=?",id).First(&user)

	if user.ID ==0 {
		return c.JSON(http.StatusNotFound,(nil))
	}
	return c.JSON(http.StatusOK,user)
  }


//   // delete user by id
//   func DeleteUserController(c echo.Context) error {
// 	// your solution here
// 	id := c.Param("id")
// 	config.DB.Delete(&model.User{},id)
// 	return c.JSON(http.StatusOK,(nil))
//   }

  // update user by id
  func UpdateUserController(c echo.Context) error {
	// your solution here
	id := c.Param("id")
	user := model.User{}

	config.DB.Where("Id = ?", id).First(&user)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound,(nil))
	}

	payload := model.User{}
	c.Bind(&payload)

	user.Name = payload.Name
	user.Email = payload.Email
	user.Password = payload.Password
	config.DB.Save(&user)

	return c.JSON(http.StatusOK,(user))
}
  
// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)
  
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}