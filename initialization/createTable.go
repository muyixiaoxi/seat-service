package initialization

import (
	"go.uber.org/zap"
	"os"
	"seat-service/model"
)

func CreateTable() {
	err := DB.AutoMigrate(
		model.Building{},
		model.Campus{},
		model.Classroom{},
		model.Reserve{},
		model.Seat{},
		model.Student{},
	)
	if err != nil {
		SeatLogger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	SeatLogger.Info("register table success")

}
