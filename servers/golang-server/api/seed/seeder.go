package seed

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tabinnorway/golang-server/api/models"
)

var locations = []models.Location{
	{
		ID:   1,
		Name: "AdO Arena",
	},
	{
		ID:   2,
		Name: "Aquarama Kristiansand",
	},
	{
		ID:   3,
		Name: "Stavanger Stupeklubb",
	},
}

var competitions = []models.Competition{
	{
		ID:         1,
		LocationId: 3,
		Title:      "Tripp Trapp 3",
		StartDate:  time.Date(2021, 9, 11, 10, 0, 0, 0, time.Local),
		EndDate:    time.Date(2021, 9, 12, 10, 0, 0, 0, time.Local),
	},
	{
		ID:         2,
		LocationId: 2,
		Title:      "Tripp Trapp 4",
		StartDate:  time.Date(2021, 11, 11, 10, 0, 0, 0, time.Local),
		EndDate:    time.Date(2021, 11, 12, 10, 0, 0, 0, time.Local),
	},
	{
		ID:         3,
		LocationId: 1,
		Title:      "Bergen Døds",
		StartDate:  time.Date(2021, 12, 5, 20, 0, 0, 0, time.Local),
		EndDate:    time.Date(2021, 12, 4, 22, 30, 0, 0, time.Local),
	},
}

var diveClubs = []models.DiveClub{
	{
		ID:        1,
		Name:      "No club registered",
		Email:     "noclub@nowhere.net",
		ContactId: 1,
	},
	{
		ID:        2,
		Name:      "Bergen Stupeklubb",
		Email:     "dagligleder@bergenstupeklubb.no",
		ContactId: 2,
	},
}

var users = []models.User{
	{
		ID:         1,
		Email:      "nobody@nowhere.net",
		Password:   "password",
		FirstName:  "Nobody",
		LastName:   "Anywhere",
		DiveClubID: 1,
	},
	{
		ID:         2,
		Email:      "andrea@bergesen.info",
		Password:   "password",
		FirstName:  "Andrea Færoøvik",
		LastName:   "Bergesen",
		DiveClubID: 2,
	},
	{
		ID:         3,
		Email:      "terje@bergesen.info",
		Password:   "password",
		FirstName:  "Terje",
		LastName:   "Bergesen",
		DiveClubID: 2,
	},
}

var dives = []models.Dive{
	{
		Title:    "Andrea dive 1",
		DiveCode: "101C",
		DiveName: "101C",
		Scores:   "6.5,7.0,7.0,6.0,6.5",
		DiverID:  1,
	},
	{
		Title:    "Andrea dive 2",
		DiveCode: "101C",
		DiveName: "101C",
		Scores:   "6.5,7.0,7.0,7.0,6.5",
		DiverID:  1,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(
		&models.Dive{},
		&models.User{},
		models.DiveClub{},
		models.Competition{},
		models.Location{},
	).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(
		models.Location{},
		models.Competition{},
		&models.User{},
		&models.Dive{},
		models.DiveClub{},
	).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Dive{}).AddForeignKey("diver_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key dive/user error: %v", err)
	}
	err = db.Debug().Model(&models.User{}).AddForeignKey("dive_club_id", "dive_clubs(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key dive/user error: %v", err)
	}
	err = db.Debug().Model(&models.Competition{}).AddForeignKey("location_id", "locations(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key dive/user error: %v", err)
	}

	for locId := range locations {
		err = db.Debug().Model(&models.Location{}).Create(&locations[locId]).Error
		if err != nil {
			log.Fatalf("cannot seed locations table: %v", err)
		}
	}

	for compId := range competitions {
		err = db.Debug().Model(&models.Competition{}).Create(&competitions[compId]).Error
		if err != nil {
			log.Fatalf("cannot seed competitions table: %v", err)
		}
	}

	for clubId := range diveClubs {
		err = db.Debug().Model(&models.DiveClub{}).Create(&diveClubs[clubId]).Error
		if err != nil {
			log.Fatalf("cannot seed diveclubs table: %v", err)
		}
	}
	for userId := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[userId]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
	for diveId := range dives {
		err = db.Debug().Model(&models.Dive{}).Create(&dives[diveId]).Error
		if err != nil {
			log.Fatalf("cannot seed dives table: %v", err)
		}
	}
}
