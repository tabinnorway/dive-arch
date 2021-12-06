package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/tabinnorway/golang-server/api/controllers"
	"github.com/tabinnorway/golang-server/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}
var diveInstance = models.Dive{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {
	var err error
	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		server.DB.LogMode(false)
		server.DB.SetLogger(nil)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(
		&models.Dive{},
		&models.User{},
		models.DiveClub{},
		models.Competition{},
		models.Location{},
	).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(
		models.Location{},
		models.Competition{},
		&models.User{},
		&models.Dive{},
		models.DiveClub{},
	).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {
	refreshUserTable()

	user := models.User{
		ID:        1,
		Email:     "andrea@bergesen.info",
		Password:  "password",
		FirstName: "Andrea",
		LastName:  "Bergesen",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {
	users := []models.User{
		{
			ID:        1,
			Email:     "andrea@bergesen.info",
			FirstName: "Andrea Færøyvik",
			LastName:  "Bergesen",
		},
		{
			ID:        2,
			Email:     "Terje@bergesen.info",
			FirstName: "Terje",
			LastName:  "Bergesen",
		},
	}

	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndDiveTable() error {
	err := server.DB.DropTableIfExists(
		&models.Dive{},
		&models.User{},
		models.DiveClub{},
		models.Competition{},
		models.Location{},
	).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(
		models.Location{},
		models.Competition{},
		&models.User{},
		&models.Dive{},
		models.DiveClub{},
	).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOneDive() (models.Dive, error) {
	err := refreshUserAndDiveTable()
	if err != nil {
		return models.Dive{}, err
	}
	user := models.User{
		ID:        1,
		Email:     "andrea@bergesen.info",
		FirstName: "Andrea Færøyvik",
		LastName:  "Bergesen",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Dive{}, err
	}
	dive := models.Dive{
		Title:    "Dive 1",
		DiverID:  1,
		DiveCode: "101C",
		DiveName: "101C",
		Scores:   "7.0,6.0,7.0,6.5,6.5",
	}
	err = server.DB.Model(&models.Dive{}).Create(&dive).Error
	if err != nil {
		return models.Dive{}, err
	}
	return dive, nil
}

func seedUsersAndDives() ([]models.User, []models.Dive, error) {
	var err error

	if err != nil {
		return []models.User{}, []models.Dive{}, err
	}
	var users = []models.User{
		{
			Email:     "andrea@bergesen.info",
			FirstName: "Andrea Færøyvik",
			LastName:  "Bergesen",
		},
	}
	var dives = []models.Dive{
		{
			Title:    "Dive 1",
			DiverID:  1,
			DiveCode: "101C",
			DiveName: "101C",
			Scores:   "7.0,6.0,7.0,6.5,6.5",
		},
		{
			Title:    "Dive 2",
			DiverID:  1,
			DiveCode: "101C",
			DiveName: "101C",
			Scores:   "7.0,6.5,7.0,7.0,6.5",
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		dives[i].DiverID = users[i].ID

		err = server.DB.Model(&models.Dive{}).Create(&dives[i]).Error
		if err != nil {
			log.Fatalf("cannot seed dives table: %v", err)
		}
	}
	return users, dives, nil
}
