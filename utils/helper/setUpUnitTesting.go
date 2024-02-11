package helper

import (
	"Todo/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// struct untuk menyambungkan ke database
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// database yang di sediakan untuk melakukan unit testing
const TEST_DB = "TODO_TEST"

// Database kosong atau tanpa table
const EMPTY_DB = "TODO_TEST_EMPTY"

func InitDBTest(dbname string) (*gorm.DB, error) {
	errE := godotenv.Load()
	if errE != nil {
		fmt.Println("Env can connect")
	}

	dbConfig := DatabaseConfig{}
	dbConfig.Host = os.Getenv("DBHOST")
	dbConfig.Port = os.Getenv("DBPORT")
	dbConfig.Username = os.Getenv("DBUSER")
	dbConfig.DBName = dbname
	dbConfig.Password = os.Getenv("DBPASS")
	// Konfigurasi koneksi database MySQL dengan GORM
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connect To Database")
	}

	return db, nil
}

func InsertDB(db *gorm.DB) error {
	tugas := []models.Tugas{
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
		{Task: "menulis", Level: "medium", Deadline: "besok", Description: "laode saady ganteng "},
	}

	err := db.Create(tugas).Error

	if err != nil {
		return err
	}

	return nil
}

func InitTestAPI() *gin.Engine {
	e := gin.Default()

	return e
}
