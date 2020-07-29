package dbmodel

type TrainResource struct {
	ID              uint    `json:"id" gorm:"primary_key"`
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}

type CreateTrainResource struct {
	DriverName      string `json:"driver_name" binding:"required"`
	OperatingStatus bool   `json:"operating_status" binding:"required"`
}

type UpdateTrainResource struct {
	DriverName      string `json:"driver_name"`
	OperatingStatus bool   `json:"operating_status"`
}