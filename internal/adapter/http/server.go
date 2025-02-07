package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"leadgen/internal/adapter/db"
	"leadgen/internal/logger"
	"leadgen/internal/usecase"
)

type Server struct {
	Router  *gin.Engine
	usecase *usecase.BuildingUsecase
	log     *zap.Logger
}

func NewServer() *Server {
	router := gin.Default()
	connect, err := db.GetDB()
	if err != nil {
		panic(err)
	}
	log, _ := logger.NewLogger()
	s := &Server{
		Router:  router,
		usecase: usecase.NewBuildingUsecase(db.NewBuildingRepository(connect), log),
		log:     log,
	}
	s.Router.Use(RecoveryMiddleware(log))
	log.Info("Starting server",
		zap.String("address", ":8080"),
	)
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
	s.Close()
}
