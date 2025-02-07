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
	u.logger.Info("Creating building",
		zap.String("name", building.Name),
		zap.String("city", building.City.Name),
		zap.Int("year_built", building.Year.Year),
		zap.Int("floor_count", building.Floor.Count),
	)
	err := u.repo.Create(building)
	if err != nil {
		u.logger.Error("Failed to create building",
			zap.String("name", building.Name),
			zap.Error(err),
		)
		return err
	}
	return err
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
