package entity

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Year struct {
	ID   int `json:"id"`
	Year int `json:"year"`
}

type Floor struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}

type Building struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	City  *City  `json:"city"`
	Year  *Year  `json:"year_built"`
	Floor *Floor `json:"floor_count"`
}
