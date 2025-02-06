package db

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

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
	var cityID int
	err := r.db.QueryRow(
		"INSERT INTO cities (name) VALUES ($1) ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name RETURNING id",
		building.City.Name,
	).Scan(&cityID)
	if err != nil {
		return err
	}

	var yearID int
	err = r.db.QueryRow(
		"INSERT INTO years (year) VALUES ($1) ON CONFLICT (year) DO UPDATE SET year=EXCLUDED.year RETURNING id",
		building.Year.Year,
	).Scan(&yearID)
	if err != nil {
		return err
	}

	var floorID int
	err = r.db.QueryRow(
		"INSERT INTO floors (count) VALUES ($1) ON CONFLICT (count) DO UPDATE SET count=EXCLUDED.count RETURNING id",
		building.Floor.Count,
	).Scan(&floorID)
	if err != nil {
		return err
	}

	query := `
        INSERT INTO buildings (name, city_id, year_id, floor_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id`
	return r.db.QueryRow(query, building.Name, cityID, yearID, floorID).Scan(&building.ID)
}

func (r *buildingRepo) GetAll(city, year, floors string) ([]entity.Building, error) {
	query := `
	SELECT 
		b.id, b.name, c.id, c.name, y.id, y.year, f.id, f.count
	FROM buildings b
	JOIN cities c ON b.city_id = c.id
	JOIN years y ON b.year_id = y.id
	JOIN floors f ON b.floor_id = f.id
	WHERE true`

	var conditions []string
	var args []interface{}

	if city != "" {
		conditions = append(conditions, "c.name = $"+strconv.Itoa(len(args)+1))
		args = append(args, city)
	}

	if year != "" {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return nil, errors.New("invalid year format")
		}
		conditions = append(conditions, "y.year = $"+strconv.Itoa(len(args)+1))
		args = append(args, yearInt)
	}

	if floors != "" {
		floorsInt, err := strconv.Atoi(floors)
		if err != nil {
			return nil, errors.New("invalid floor count format")
		}
		conditions = append(conditions, "f.count = $"+strconv.Itoa(len(args)+1))
		args = append(args, floorsInt)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []entity.Building
	for rows.Next() {
		var building entity.Building
		var cityID int
		var cityName string
		var yearID, yearBuilt int
		var floorID, floorCount int

		err := rows.Scan(
			&building.ID,
			&building.Name,
			&cityID,
			&cityName,
			&yearID,
			&yearBuilt,
			&floorID,
			&floorCount,
		)
		if err != nil {
			return nil, err
		}

		building.City = &entity.City{ID: cityID, Name: cityName}
		building.Year = &entity.Year{ID: yearID, Year: yearBuilt}
		building.Floor = &entity.Floor{ID: floorID, Count: floorCount}

		buildings = append(buildings, building)
	}

	return buildings, nil
}
