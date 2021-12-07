package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Location struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Address    string    `gorm:"size:100;not null" json:"address"`
	Postalcode string    `gorm:"size:100;not null" json:"postalcode"`
	Country    string    `gorm:"size:100;not null" json:"country"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (dc *Location) Prepare() {
	dc.ID = 0
	dc.Name = html.EscapeString(strings.TrimSpace(dc.Name))
	dc.Name = html.EscapeString(strings.TrimSpace(dc.Address))
	dc.Name = html.EscapeString(strings.TrimSpace(dc.Postalcode))
	dc.Name = html.EscapeString(strings.TrimSpace(dc.Country))
	dc.CreatedAt = time.Now()
	dc.UpdatedAt = time.Now()
}

func (dc *Location) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if dc.Name == "" {
			return errors.New("required club name")
		}
		return nil

	default:
		if dc.Name == "" {
			return errors.New("required name")
		}
		return nil
	}
}

func (dc *Location) SaveLocation(db *gorm.DB) (*Location, error) {
	err := db.Create(&dc).Error
	if err != nil {
		return &Location{}, err
	}
	return dc, nil
}

func (dc *Location) FindAllLocations(db *gorm.DB) (*[]Location, error) {
	ret := []Location{}
	err := db.Model(&Location{}).Limit(1000).Find(&ret).Error
	if err != nil {
		return &[]Location{}, err
	}
	return &ret, err
}

func (dc *Location) FindLocationByID(db *gorm.DB, uid uint32) (*Location, error) {
	err := db.Model(Location{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &Location{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Location{}, errors.New("Location Not Found")
	}
	return dc, err
}

func (dc *Location) UpdateALocation(db *gorm.DB, uid uint32) (*Location, error) {
	db = db.Model(&Location{}).Where("id = ?", uid).Take(&Location{}).UpdateColumns(
		map[string]interface{}{
			"name":       dc.Name,
			"address":    dc.Address,
			"country":    dc.Country,
			"postalcode": dc.Postalcode,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Location{}, db.Error
	}
	// This is the display the updated user
	err := db.Model(&Location{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &Location{}, err
	}
	return dc, nil
}

func (dc *Location) DeleteALocation(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Model(&Location{}).Where("id = ?", uid).Take(&Location{}).Delete(&Location{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
