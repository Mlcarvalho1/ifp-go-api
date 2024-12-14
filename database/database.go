package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"ifp-analysis.com/config"
)

var DB *gorm.DB // Variável global para armazenar a instância do banco

func ConnectDb() {
	dsn := "host=" + config.GetEnv("DB_HOST") +
		" user=" + config.GetEnv("DB_USER") +
		" password=" + config.GetEnv("DB_PASSWORD") +
		" dbname=" + config.GetEnv("DB_NAME") +
		" port=" + config.GetEnv("DB_PORT")

	// Inicializar a conexão com o banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	// Atribuir a conexão à variável global DB
	DB = db

	log.Println("Connected to the database")
	log.Println("Running migrations")

	// Adicione migrações, se necessário
	// db.AutoMigrate(&models.User{}) // Substitua pelo modelo relevante
}
