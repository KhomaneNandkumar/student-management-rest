package initializers

import (
	"fmt"

	"example.com/student-management/internal/models"
	"github.com/projectdiscovery/gologger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DB_DRIVER = "postgres"

var DB *gorm.DB
var err error

func generateDBURL(c Config) string {
	if c.DBHOST == "" || c.DBPORT == "" || c.DBUSER == "" || c.DBPASSWORD == "" || c.DBNAME == "" {
		gologger.Fatal().Msg("Database configuration is missing required fields. Check environment variables.")
	}

	// Debug: Print generated DB Connection String (excluding password)
	fmt.Println("Generated DB Connection String:", fmt.Sprintf(
		"host=%s user=%s password=******** dbname=%s port=%s sslmode=disable",
		c.DBHOST, c.DBUSER, c.DBNAME, c.DBPORT,
	))

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.DBHOST, c.DBUSER, c.DBPASSWORD, c.DBNAME, c.DBPORT,
	)
}

func ConnectDB() error {
	c := GetConfig()
	dbURL := generateDBURL(c)

	// Use the correct DB URL format
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		gologger.Fatal().Msgf("Error Connecting to Database: %s\n", err)
		return err
	}

	// Run AutoMigrate
	DB.AutoMigrate(&models.Student{})
	gologger.Info().Msgf("Connected To The Database: %s", c.DBNAME)
	return nil
}
