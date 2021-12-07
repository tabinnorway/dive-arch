package controllertests

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
var postInstance = models.Dive{}

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
		models.DiveClub{},
		models.Competition{},
		&models.User{},
		&models.Dive{},
	).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Email:    "pet@gmail.com",
		Password: "password",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func seedUsers() ([]models.User, error) {
	var err error
	if err != nil {
		return nil, err
	}
	users := []models.User{
		{
			Email:    "steven@gmail.com",
			Password: "password",
		},
		{
			Email:    "kenny@gmail.com",
			Password: "password",
		},
	}
	for i := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func refreshUserAndDiveTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Dive{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Dive{}).Error
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
		Email:    "sam@gmail.com",
		Password: "password",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Dive{}, err
	}
	post := models.Dive{
		Title:   "This is the title sam",
		DiverID: user.ID,
	}
	err = server.DB.Model(&models.Dive{}).Create(&post).Error
	if err != nil {
		return models.Dive{}, err
	}
	return post, nil
}

func seedUsersAndDives() ([]models.User, []models.Dive, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Dive{}, err
	}
	var users = []models.User{
		{
			Email:    "steven@gmail.com",
			Password: "password",
		},
		{
			Email:    "magu@gmail.com",
			Password: "password",
		},
	}
	var posts = []models.Dive{
		{
			Title: "Title 1",
		},
		{
			Title: "Title 2",
		},
	}

	for i := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].DiverID = users[i].ID

		err = server.DB.Model(&models.Dive{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
	return users, posts, nil
}
