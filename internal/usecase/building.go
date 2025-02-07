package usecase

import (
	"fmt"
	"go.uber.org/zap"
	"leadgen/internal/adapter/db"
	"leadgen/internal/entity"
)

type BuildingUsecase struct {
	repo   db.BuildingRepository
	logger *zap.Logger
}

func NewBuildingUsecase(repo db.BuildingRepository, logger *zap.Logger) *BuildingUsecase {
	return &BuildingUsecase{repo: repo, logger: logger}
}

func (u *BuildingUsecase) Create(building *entity.Building) error {
	cityName := "N/A"
	if building.City != nil {
		cityName = building.City.Name
	}
	yearBuilt := 0
	if building.Year != nil {
		yearBuilt = building.Year.Year
	}
	floorCount := 0
	if building.Floor != nil {
		floorCount = building.Floor.Count
	}

	u.logger.Info("Creating building",
		zap.String("name", building.Name),
		zap.String("city", cityName),
		zap.Int("year_built", yearBuilt),
		zap.Int("floor_count", floorCount),
	)
	err := u.repo.Create(building)
	if err != nil {
		u.logger.Error("Failed to create building",
			zap.String("name", building.Name),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (u *BuildingUsecase) GetAll(city, year, floors string) ([]entity.Building, error) {
	buildings, err := u.repo.GetAll(city, year, floors)

	if err != nil {
		u.logger.Error("Failed to fetch buildings",
			zap.String("filters", fmt.Sprintf("city=%s, year=%s, floors=%s", city, year, floors)),
			zap.Error(err),
		)
		return nil, err
	}
	return buildings, nil
}
