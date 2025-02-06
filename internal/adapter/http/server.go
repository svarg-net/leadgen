package http

import (
	"leadgen/internal/adapter/db"
	"leadgen/internal/entity"
	"leadgen/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	usecase *usecase.BuildingUsecase
}

func NewServer() *Server {
	router := gin.Default()
	connect, err := db.GetDB()
	if err != nil {
		panic(err)
	}
	s := &Server{
		Router:  router,
		usecase: usecase.NewBuildingUsecase(db.NewBuildingRepository(connect)),
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.Router.POST("/buildings", s.createBuilding)
	s.Router.GET("/buildings", s.getBuildings)
}

func (s *Server) createBuilding(c *gin.Context) {
	var input struct {
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Преобразуем входные данные в структуру Building
	building := &entity.Building{}

	// Создаем здание
	if err := s.usecase.Create(building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем созданное здание
	c.JSON(http.StatusCreated, building)
}

func (s *Server) getBuildings(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year_built")
	floors := c.Query("floor_count")

	buildings, err := s.usecase.GetAll(city, year, floors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, buildings)
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

func (s *Server) Close() {
	// Закрытие ресурсов, если необходимо
}
