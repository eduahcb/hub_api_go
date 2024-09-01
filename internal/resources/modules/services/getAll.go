package modulesservices

import (
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/modules"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
)

func GeAll(db *database.Database) ([]modules.ModuleResponse, error) {
  var modulesEntity []entities.Module
  modulesResponse := []modules.ModuleResponse{}
  
  err := db.Client.Find(&modulesEntity).Error

  if err != nil {
    return modulesResponse, &customErrors.InternalServerError{}
  }

  for _, module := range modulesEntity {
    modulesResponse = append(modulesResponse, modules.ModuleResponse{
      ID: module.ID,
      Name: module.Name,
    })
  }

  return modulesResponse, nil
}
