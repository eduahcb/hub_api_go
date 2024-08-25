package main

import (
	"fmt"
	"log"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/pkg/security"
)

func init() {
	entities := []interface{}{
		&entities.User{},
		&entities.Tech{},
		&entities.Module{},
		&entities.Level{},
	}

	config.SetEntities(entities)

	err := config.Init()

	if err != nil {
		log.Fatal("something went wrong!!! ", err)
	}
}

func main() {
	client := config.GetDBClient()

	modules := []entities.Module{
		{Name: "1 - Iniciando no frontend"},
		{Name: "2 - Iniciando no backend"},
	}

	err := client.Create(modules).Error

	if err != nil {
		fmt.Println("error to create modules")
	}

	levels := []*entities.Level{
		{Name: "Iniciante"},
		{Name: "Intermediário"},
		{Name: "Avançado"},
	}

	err = client.Create(levels).Error

	if err != nil {
		fmt.Println("error to create levels")
	}

	hashPassword, _ := security.CreateHashPassword("12345678")

	user := entities.User{
		Name:     "Crispim",
		Email:    "crispim@test.com",
		Password: string(hashPassword),
		ModuleID: 1,
	}

	err = client.Create(&user).Error

	if err != nil {
		fmt.Println("error to create user")
	}

	fmt.Println("seeding success!!!")
}
