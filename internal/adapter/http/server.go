package http

import (
	"leadgen/internal/adapter/db"
	"leadgen/internal/usecase"

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

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

func (s *Server) Close() {
	// Закрытие ресурсов, если необходимо
}
