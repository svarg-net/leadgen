package http

import (
	"leadgen/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new building
// @Description Creates a new building in the database
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body entity.Building true "Building details"
// @Success 201 {object} entity.Building
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /buildings [post]
func (s *Server) createBuilding(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		CityName string `json:"city" binding:"required"`
		Year     int    `json:"year_built" binding:"required"`
		Floor    int    `json:"floor_count" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	building := &entity.Building{
		Name:  input.Name,
		City:  &entity.City{Name: input.CityName},
		Year:  &entity.Year{Year: input.Year},
		Floor: &entity.Floor{Count: input.Floor},
	}

	if err := s.usecase.Create(building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
