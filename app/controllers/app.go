package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"net/http"
	"revel-rest/dbmodel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func init() {
	fmt.Println("Connecting to DB")
	dbmodel.ConnectDataBase()
}

// GetTrain handles GET on train resource
func (c App) GetTrain() revel.Result {

	//var train dbmodel.TrainResource
	var trains []dbmodel.TrainResource
	dbmodel.DB.Find(&trains)

	c.Response.Status = http.StatusOK
	return c.RenderJSON(trains)
}

// GetTrain handles GET on train resource
func (c App) GetTrainByID() revel.Result {
	fmt.Println("ID: ", c.Params.Route.Get("train-id"))
	var train dbmodel.TrainResource
	if err := dbmodel.DB.Where("id = ?", c.Params.Route.Get("train-id")).First(&train).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(train)
	}

	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

// CreateTrain handles POST on train resource
func (c App) CreateTrain() revel.Result {
	var input dbmodel.CreateTrainResource

	if err := c.Params.BindJSON(&input); err != nil {
		c.Response.Status = http.StatusCreated
		return c.RenderJSON(input)
	}

	fmt.Println(">>>>>>>>>>>>>: ", input.DriverName, input.OperatingStatus)
	// Create Train
	train := dbmodel.TrainResource{DriverName: input.DriverName, OperatingStatus: input.OperatingStatus}
	dbmodel.DB.Create(&train)

	c.Response.Status = http.StatusCreated
	return c.RenderJSON(input)
}

// RemoveTrain implements DELETE on train resource
func (c App) RemoveTrain() revel.Result {
	fmt.Println("ID: ", c.Params.Route.Get("train-id"))
	var train dbmodel.TrainResource
	if err := dbmodel.DB.Where("id = ?", c.Params.Route.Get("train-id")).First(&train).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(train)
	}

	dbmodel.DB.Delete(&train)
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

// RemoveTrain implements DELETE on train resource
func (c App) UpdateTrain() revel.Result {
	fmt.Println("ID: ", c.Params.Route.Get("train-id"))

	var train dbmodel.TrainResource
	if err := dbmodel.DB.Where("id = ?", c.Params.Route.Get("train-id")).First(&train).Error; err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(train)
	}

	// Validate input
	var input dbmodel.UpdateTrainResource
	if err := c.Params.BindJSON(&input); err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(train)
	}

	dbmodel.DB.Model(&train).Update(&input)
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}