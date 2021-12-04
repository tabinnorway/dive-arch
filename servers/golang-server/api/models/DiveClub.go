package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type DiveClub struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null" json:"name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	ContactId uint32 `gorm:"not null" json:"contact_id"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (dc *DiveClub) Prepare() {
	dc.ID = 0
	dc.Email = html.EscapeString(strings.TrimSpace(dc.Email))
	dc.Name = html.EscapeString(strings.TrimSpace(dc.Name))
	dc.CreatedAt = time.Now()
	dc.UpdatedAt = time.Now()
}

func (dc *DiveClub) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if dc.Name == "" {
			return errors.New("required club name")
		}
		if dc.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(dc.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if dc.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(dc.Email); err != nil {
			return errors.New("invalid email")
		}
		if dc.Name == "" {
			return errors.New("required name")
		}
		return nil
	}
}

func (dc *DiveClub) SaveDiveClub(db *gorm.DB) (*DiveClub, error) {

	err := db.Debug().Create(&dc).Error
	if err != nil {
		return &DiveClub{}, err
	}
	return dc, nil
}

func (dc *DiveClub) FindAllDiveClubs(db *gorm.DB) (*[]DiveClub, error) {
	ret := []DiveClub{}
	err := db.Debug().Model(&DiveClub{}).Limit(1000).Find(&ret).Error
	if err != nil {
		return &[]DiveClub{}, err
	}
	return &ret, err
}

func (dc *DiveClub) FindDiveClubByID(db *gorm.DB, uid uint32) (*DiveClub, error) {
	err := db.Debug().Model(DiveClub{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &DiveClub{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &DiveClub{}, errors.New("DiveClub Not Found")
	}
	return dc, err
}

func (dc *DiveClub) UpdateADiveClub(db *gorm.DB, uid uint32) (*DiveClub, error) {
	// To hash the password
	db = db.Debug().Model(&DiveClub{}).Where("id = ?", uid).Take(&DiveClub{}).UpdateColumns(
		map[string]interface{}{
			"name":       dc.Name,
			"email":      dc.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &DiveClub{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&DiveClub{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &DiveClub{}, err
	}
	return dc, nil
}

func (dc *DiveClub) DeleteADiveClub(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&DiveClub{}).Where("id = ?", uid).Take(&DiveClub{}).Delete(&DiveClub{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
