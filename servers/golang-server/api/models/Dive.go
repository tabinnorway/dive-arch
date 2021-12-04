package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Dive struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title         string    `gorm:"size:255;not null;unique" json:"title"`
	Diver         User      `json:"diver"`
	DiverID       uint64    `gorm:"not null" json:"diver_id"`
	DiveCode      string    `gorm:"size:255;not null" json:"dive_code"`
	DiveName      string    `gorm:"not null" json:"dive_name"`
	Scores        string    `gorm:"not null" json:"scores"`
	MovieLocation string    `gorm:"not null" json:"movie_location"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (d *Dive) Prepare() {
	d.ID = 0
	d.Title = html.EscapeString(strings.TrimSpace(d.Title))
	d.Title = html.EscapeString(strings.TrimSpace(d.Scores))
	d.MovieLocation = html.EscapeString(strings.TrimSpace(d.MovieLocation))
	d.Diver = User{}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

func (d *Dive) Validate() error {

	if d.Title == "" {
		return errors.New("required title")
	}
	if d.DiverID < 1 {
		return errors.New("required diver")
	}
	return nil
}

func (d *Dive) SaveDive(db *gorm.DB) (*Dive, error) {
	err := db.Debug().Model(&Dive{}).Create(&d).Error
	if err != nil {
		return &Dive{}, err
	}
	if d.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", d.DiverID).Take(&d.Diver).Error
		if err != nil {
			return &Dive{}, err
		}
	}
	return d, nil
}

func (d *Dive) FindAllDives(db *gorm.DB) (*[]Dive, error) {
	dives := []Dive{}
	err := db.Debug().Model(&Dive{}).Limit(1000).Find(&dives).Error
	if err != nil {
		return &[]Dive{}, err
	}
	if len(dives) > 0 {
		for i := range dives {
			err := db.Debug().Model(&User{}).Where("id = ?", dives[i].DiverID).Take(&dives[i].Diver).Error
			if err != nil {
				return &[]Dive{}, err
			}
		}
	}
	return &dives, nil
}

func (d *Dive) FindDiveByID(db *gorm.DB, pid uint64) (*Dive, error) {
	err := db.Debug().Model(&Dive{}).Where("id = ?", pid).Take(&d).Error
	if err != nil {
		return &Dive{}, err
	}
	if d.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", d.DiverID).Take(&d.Diver).Error
		if err != nil {
			return &Dive{}, err
		}
	}
	return d, nil
}

func (d *Dive) UpdateADive(db *gorm.DB) (*Dive, error) {
	err := db.Debug().Model(&Dive{}).Where("id = ?", d.ID).Updates(Dive{Title: d.Title, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Dive{}, err
	}
	if d.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", d.DiverID).Take(&d.Diver).Error
		if err != nil {
			return &Dive{}, err
		}
	}
	return d, nil
}

func (d *Dive) DeleteADive(db *gorm.DB, pid uint64, uid uint64) (int64, error) {

	db = db.Debug().Model(&Dive{}).Where("id = ? and diver_id = ?", pid, uid).Take(&Dive{}).Delete(&Dive{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("dive not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
