package http

import (
	"go.uber.org/zap"
	"leadgen/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new building
// @Description Creates a new building in the database
// @Tags buildings
// @Accept json
// @Produce json
//
//	@Param building body struct {
//	    Name     string `json:"name" binding:"required"`
//	    CityName string `json:"city" binding:"required"`
//	    Year     int    `json:"year_built" binding:"required"`
//	    Floor    int    `json:"floor_count" binding:"required"`
//	} true "Building details"
//
// @Success 201 {object} entity.Building
// @Failure 400 {object} map[string]string "Invalid input or missing fields"
// @Failure 409 {object} map[string]string "Building already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /buildings [post]
func (s *Server) createBuilding(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		CityName string `json:"city" binding:"required"`
		Year     int    `json:"year_built" binding:"required"`
		Floor    int    `json:"floor_count" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		s.log.Warn("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.Name == "" || input.CityName == "" || input.Year <= 0 || input.Floor <= 0 {
		s.log.Warn("Missing or invalid fields in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required and must be valid"})
		return
	}

	building := &entity.Building{
		Name:  input.Name,
		City:  &entity.City{Name: input.CityName},
		Year:  &entity.Year{Year: input.Year},
		Floor: &entity.Floor{Count: input.Floor},
	}

	if err := s.usecase.Create(building); err != nil {
		s.log.Error("Failed to create building", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, building)
}

// @Summary Get all buildings
// @Description Retrieves a list of buildings based on optional filters
// @Tags buildings
// @Accept json
// @Produce json
// @Param city query string false "City name"
// @Param year_built query string false "Year built"
// @Param floor_count query string false "Number of floors"
// @Success 200 {array} entity.Building
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /buildings [get]
func (s *Server) getBuildings(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year_built")
	floors := c.Query("floor_count")

	buildings, err := s.usecase.GetAll(city, year, floors)
	if err != nil {
		s.log.Error("Failed to get building",
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, buildings)
}
