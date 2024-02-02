package config

import (
	"Todo/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func InitializeDatabase() *gorm.DB {

	errE := godotenv.Load()
	if errE != nil {
		fmt.Println("Env can connect")
	}

	dbConfig := DatabaseConfig{}
	dbConfig.Host = os.Getenv("DBHOST")
	dbConfig.Port = os.Getenv("DBPORT")
	dbConfig.Username = os.Getenv("DBUSER")
	dbConfig.DBName = os.Getenv("DBNAME")
	dbConfig.Password = os.Getenv("DBPASS")
	// Konfigurasi koneksi database MySQL dengan GORM
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect To Database")
	}

	// // Migrate tabel pengguna
	db.AutoMigrate(&models.Tugas{})
	db.AutoMigrate(&models.Notification{})

	return db
}
