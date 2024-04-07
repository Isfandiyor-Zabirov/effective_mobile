package handlers

import (
	"effective_mobile_tech_task/internal/entitites/cars"
	"effective_mobile_tech_task/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCars - Сохранение машины в базу по госномеру
// @Summary Сохранение машины в базу по госномеру
// @ID create-cars
// @Tags Машины
// @Produce json
// @Param id body cars.AddCarsRequest true "Госномера машин"
// @Success 200 {string} string "Create successful"
// @Failure 400 {string} string "Something went wrong"
// @Router /api/v1/cars [post]
func CreateCars(c *gin.Context) {
	var (
		request cars.AddCarsRequest
		err     error
	)

	fmt.Println("checkpoint 1")

	if err = c.ShouldBindJSON(&request); err != nil {
		logger.Error.Printf("CreateCars handler cannot bind the request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Wrong request format"})
		return
	}

	if err = cars.GetAndSaveCars(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cars successfully saved"})
}

// UpdateCar - Обновление данных машины в базе
// @Summary Обновление данных машины в базе
// @ID update-cars
// @Tags Машины
// @Produce json
// @Param id body cars.Car true "Данные машины"
// @Success 200 {string} string "Update successful"
// @Failure 400 {string} string "Something went wrong"
// @Router /api/v1/cars [put]
func UpdateCar(c *gin.Context) {
	var (
		request cars.Car
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		logger.Error.Println("UpdateCar handler cannot bind the request:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Wrong request format"})
		return
	}

	if err = cars.UpdateCar(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car successfully updated"})
}

// DeleteCar - Удаление машины из базы по идентификатору (ID)
// @Summary Удаление машины из базы по идентификатору (ID)
// @ID delete-cars
// @Tags Машины
// @Produce json
// @Param id path string true "ID машины"
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {string} string "Something went wrong"
// @Router /api/v1/cars [delete]
func DeleteCar(c *gin.Context) {
	var (
		carIdStr = c.Param("id")
		carIdInt = 0
		err      error
	)

	carIdInt, err = strconv.Atoi(carIdStr)
	if err != nil {
		logger.Error.Println("DeleteCar handler cannot convert car id from params:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Wrong car id in params"})
		return
	}

	if err = cars.DeleteCar(carIdInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car successfully deleted"})
}

// GetCars - Получение данных машин по фильтру с пагинацией и фильтрацией.
// @Summary Получение данных машин по фильтру с пагинацией и фильтрацией.
// @ID get-cars
// @Tags Машины
// @Produce json
// @Param regNum 	 			query string   false "Госномер машины"
// @Param mark 	 	        	query string   false "Марка машины"
// @Param model 	    		query string   false "Модел машины"
// @Param year 	 				query integer  false "Год производства машины"
// @Param ownerID  		 		query integer  false "ID владельца машины"
// @Param page 	 		 		query integer  false "Страница"
// @Param page_limit 	 		query integer  false "Количество рядов на странице(для пагинации)"
// @Success 200 {object} cars.GetCarsResponse
// @Failure 400 {string} string  "Something went wrong"
// @Router /api/v1/cars [get]
func GetCars(c *gin.Context) {
	var (
		filter cars.GetCarsFilter
	)

	if err := c.Bind(&filter); err != nil {
		logger.Error.Println("GetCars handler cannot bind filters:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Wrong filter struct"})
		return
	}

	resp, err := cars.GetCars(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
