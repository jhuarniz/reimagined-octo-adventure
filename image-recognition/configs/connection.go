package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to the database
func DBConnection() {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=UTC", EnvDBHost(), EnvDBUser(), EnvDBPass(), EnvDBName(), EnvDBPort())
	var err error
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("DB Connected...")
}
