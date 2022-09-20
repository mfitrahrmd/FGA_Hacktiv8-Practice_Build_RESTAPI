package controllers

import (
	"FGA_Hacktiv8-Practice_Build_RESTAPI/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type InDB struct {
	DB *gorm.DB
}

func NewPostgresDB() *InDB {
	db, err := gorm.Open(postgres.Open("postgresql://developer:developer@localhost:5432/fga_hacktiv8_sesi7?sslmode=disable"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &InDB{
		DB: db,
	}
}

func (idb *InDB) Migrate() {
	err := idb.DB.AutoMigrate(structs.Person{})
	if err != nil {
		log.Fatalln(err.Error())
	}
}
