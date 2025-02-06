package db

import (
	"database/sql"

	"leadgen/internal/entity"
)

type BuildingRepository interface {
	Create(*entity.Building) error
	GetAll(city, year, floors string) ([]entity.Building, error)
}

type buildingRepo struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) BuildingRepository {
	return &buildingRepo{db: db}
}

func (r *buildingRepo) Create(building *entity.Building) error {
	return nil
}

func (r *buildingRepo) GetAll(city, year, floors string) ([]entity.Building, error) {

	return []entity.Building{}, nil
}
