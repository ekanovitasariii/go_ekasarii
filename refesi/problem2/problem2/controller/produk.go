package controller

import (
	"net/http"
	"problem2/config"
	"problem2/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var logs []model.LogProduk

	if err := config.DB.Find(&logs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK,logs)
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	log_product := model.LogProduk{}

	config.DB.Where("ID = ?", id).First(&log_product)

	if log_product.ID == 0 {
		return c.JSON(http.StatusNotFound,nil)
	}

	return c.JSON(http.StatusOK,log_product)
}

func CreateBookController(c echo.Context) error {
	log_product := model.LogProduk{}
	c.Bind(&log_product)

	if err := config.DB.Save(&log_product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, log_product)
}

// func DeleteBookController(c echo.Context) error {
// 	id := c.Param("id")

// 	config.DB.Delete(&model.Book{}, id)

// 	return c.JSON(http.StatusOK, nil)
// }

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	log_product := model.LogProduk{}

	config.DB.Where("ID = ?", id).First(&log_product)

	if log_product.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	payload := model.LogProduk{}
	c.Bind(&payload)

	log_product.ID = payload.ID
	log_product.Status = payload.Status
	log_product.TotalQty = payload.TotalQty
	log_product.TotalPrice= payload.TotalPrice
	config.DB.Save(&log_product)

	return c.JSON(http.StatusOK,log_product)
}