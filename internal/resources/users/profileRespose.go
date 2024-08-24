package users

type ModuleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProfileResponse struct {
	ID      uint           `json:"id"`
	Name    string         `json:"name"`
	Contact string         `json:"contact"`
	Bio     string         `json:"bio"`
	Module  ModuleResponse `json:"module"`
}
