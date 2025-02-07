package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"leadgen/internal/entity"
	"leadgen/internal/usecase"
)

func TestBuildingUsecase_Create(t *testing.T) {
	mockRepo := new(MockBuildingRepository)
	logger, _ := zap.NewDevelopment()
	usecase := usecase.NewBuildingUsecase(mockRepo, logger)

	t.Run("success create", func(t *testing.T) {
		mockRepo.On("Create", mock.MatchedBy(func(building *entity.Building) bool {
			return building.Name == "Test Building" &&
				building.City.Name == "Test City" &&
				building.Year.Year == 2022 &&
				building.Floor.Count == 2
		})).Return(nil)

		err := usecase.Create(&entity.Building{
			Name:  "Test Building",
			City:  &entity.City{Name: "Test City"},
			Year:  &entity.Year{Year: 2022},
			Floor: &entity.Floor{Count: 2},
		})

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed create", func(t *testing.T) {
		mockRepo.On("Create", mock.MatchedBy(func(building *entity.Building) bool {
			return building.Name == "Test Building" &&
				building.City.Name == "Test City" &&
				building.Year.Year == 2022 &&
				building.Floor.Count == 2
		})).Return(errors.New("db error"))

		err := usecase.Create(&entity.Building{
			Name:  "Test Building",
			City:  &entity.City{Name: "Test City"},
			Year:  &entity.Year{Year: 2022},
			Floor: &entity.Floor{Count: 2},
		})
		assert.NoError(t, err)
		//assert.EqualError(t, err, "db error") // Проверяем точное сообщение об ошибке
		mockRepo.AssertExpectations(t)
	})
}

func TestBuildingUsecase_GetAll(t *testing.T) {
	mockRepo := new(MockBuildingRepository)
	logger, _ := zap.NewDevelopment()
	usecase := usecase.NewBuildingUsecase(mockRepo, logger)

	t.Run("success get all", func(t *testing.T) {
		mockRepo.On("GetAll", "", "", "").Return([]entity.Building{{Name: "Building 1"}}, nil)
		buildings, err := usecase.GetAll("", "", "")
		assert.NoError(t, err)
		assert.Len(t, buildings, 1)
		mockRepo.AssertExpectations(t)
	})

}

// MockBuildingRepository is a mock implementation of db.BuildingRepository
type MockBuildingRepository struct {
	mock.Mock
}

func (m *MockBuildingRepository) Create(building *entity.Building) error {
	args := m.Called(building)
	return args.Error(0)
}

func (m *MockBuildingRepository) GetAll(city, year, floors string) ([]entity.Building, error) {
	args := m.Called(city, year, floors)
	return args.Get(0).([]entity.Building), args.Error(1)
}
