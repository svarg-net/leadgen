package test

import (
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"

	"leadgen/internal/entity"
	"leadgen/internal/usecase"
)

func BenchmarkBuildingUsecase_Create(b *testing.B) {
	mockRepo := new(MockBuildingRepository)
	logger, _ := zap.NewDevelopment()
	usecase := usecase.NewBuildingUsecase(mockRepo, logger)

	mockRepo.On("Create", mock.Anything).Return(nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := usecase.Create(&entity.Building{Name: "Benchmark Building"})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildingUsecase_GetAll(b *testing.B) {
	mockRepo := new(MockBuildingRepository)
	logger, _ := zap.NewDevelopment()
	usecase := usecase.NewBuildingUsecase(mockRepo, logger)

	mockRepo.On("GetAll", "", "", "").Return([]entity.Building{{Name: "Building 1"}}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := usecase.GetAll("", "", "")
		if err != nil {
			b.Fatal(err)
		}
	}
}
