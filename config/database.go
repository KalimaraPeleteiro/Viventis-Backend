package config

import (
	"fmt"
	"log"
	"os"
	"viventis/schemas"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	_ = godotenv.Load()
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("SSL_MODE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro na conexão ao banco de dados: ", err)
	}

	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		log.Fatal("Erro ao atualizar o banco de dados: ", err)
	}

	fmt.Println("Conexão com o banco bem-sucedida.")
}
