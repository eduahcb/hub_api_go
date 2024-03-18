package techs

type LevelResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TechResponse struct {
	ID    uint          `json:"id"`
	Name  string        `json:"name"`
	Level LevelResponse `json:"level"`
}
