package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding $HOME directory")	
	}
	_ = os.Mkdir(fmt.Sprintf("%s/pokecommit/", home), os.ModePerm)
	conn, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/pokecommit/pokecommit.db", home)), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database")
	}
	db = conn;
	db.AutoMigrate(&PokemonModel{})
}

func GetDB() *gorm.DB {
	return db;
}