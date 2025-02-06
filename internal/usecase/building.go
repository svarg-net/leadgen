package usecase

import (
	"leadgen/internal/adapter/db"
	"leadgen/internal/entity"
)

type BuildingUsecase struct {
	repo db.BuildingRepository
}

func NewBuildingUsecase(repo db.BuildingRepository) *BuildingUsecase {
	return &BuildingUsecase{repo: repo}
}

func (u *BuildingUsecase) Create(building *entity.Building) error {
	return u.repo.Create(building)
}

func (u *BuildingUsecase) GetAll(city, year, floors string) ([]entity.Building, error) {
	return u.repo.GetAll(city, year, floors)
}
