package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Competition struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title      string    `gorm:"size:255;not null;unique" json:"title"`
	Location   Location  `json:"location"`
	LocationId uint64    `gorm:"not null" json:"location_id"`
	StartDate  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"start_date"`
	EndDate    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"end_date"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Competition) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Competition) Validate() error {
	if p.Title == "" {
		return errors.New("required title")
	}
	return nil
}

func (p *Competition) SaveCompetition(db *gorm.DB) (*Competition, error) {
	err := db.Model(&Competition{}).Create(&p).Error
	if err != nil {
		return &Competition{}, err
	}
	return p, nil
}

func (p *Competition) FindAllCompetitions(db *gorm.DB) (*[]Competition, error) {
	ret := []Competition{}
	err := db.Model(&Competition{}).Limit(1000).Find(&ret).Error
	if err != nil {
		return &[]Competition{}, err
	}
	return &ret, nil
}

func (c *Competition) FindCompetitionByID(db *gorm.DB, pid uint64) (*Competition, error) {
	err := db.Model(&Competition{}).Where("id = ?", pid).Take(&c).Error
	if err != nil {
		return &Competition{}, err
	}
	return c, nil
}

func (c *Competition) UpdateACompetition(db *gorm.DB) (*Competition, error) {
	err := db.Model(&Competition{}).Where("id = ?", c.ID).Updates(Competition{Title: c.Title, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Competition{}, err
	}
	return c, nil
}

func (p *Competition) DeleteACompetition(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Model(&Competition{}).Where("id = ? and diver_id = ?", pid, uid).Take(&Competition{}).Delete(&Competition{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("dive not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
