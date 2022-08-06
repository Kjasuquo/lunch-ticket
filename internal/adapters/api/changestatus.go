package api

import (
	"github.com/decadevs/lunch-api/internal/core/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (u *HTTPHandler) UpdateBrunchFoodStatus(c *gin.Context) {
	user, err := u.GetKitchenStaffFromContext(c)
	log.Println("user in context", user)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, month, day := time.Now().Date()

	food, err := u.UserService.FindBrunchByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	type StatusUpdate struct {
		Status string `json:"status"`
	}

	var status StatusUpdate

	if err := c.BindJSON(&status); err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	errS := u.UserService.UpdateStatus(food, status.Status)
	if errS != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"error updating food"})
		return
	}

	currentFood, _ := u.UserService.FindBrunchByDate(year, month, day)
	helpers.JSON(c, "food status updated successfully", http.StatusOK, currentFood, nil)
}

func (u *HTTPHandler) UpdateDinnerFoodStatus(c *gin.Context) {
	user, err := u.GetKitchenStaffFromContext(c)
	log.Println("user in context", user)
	if err != nil {
		helpers.JSON(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}
	year, month, day := time.Now().Date()

	food, err := u.UserService.FindDinnerByDate(year, month, day)
	if err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	type StatusUpdate struct {
		Status string `json:"status"`
	}

	var status StatusUpdate

	if err := c.BindJSON(&status); err != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"internal server error"})
		return
	}

	errS := u.UserService.UpdateStatus(food, status.Status)
	if errS != nil {
		helpers.JSON(c, "internal server error", 500, nil, []string{"error updating food"})
		return
	}

	currentFood, _ := u.UserService.FindDinnerByDate(year, month, day)
	helpers.JSON(c, "food status updated successfully", http.StatusOK, currentFood, nil)
}