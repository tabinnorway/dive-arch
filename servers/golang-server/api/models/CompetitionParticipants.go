package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CompetitionParticipant struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Diver      User      `json:"diver"`
	DiverID    uint64    `gorm:"not null" json:"diver_id"`
	Location   Location  `json:"location"`
	LocationID uint64    `gorm:"not null" json:"location_id"`
	Comment    string    `gorm:"size:1024" json:"name"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (dc *CompetitionParticipant) Prepare() {
	dc.ID = 0
	dc.Comment = html.EscapeString(strings.TrimSpace(dc.Comment))
	dc.CreatedAt = time.Now()
	dc.UpdatedAt = time.Now()
}

func (dc *CompetitionParticipant) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		return nil

	default:
		return nil
	}
}

func (dc *CompetitionParticipant) SaveCompetitionParticipant(db *gorm.DB) (*CompetitionParticipant, error) {
	err := db.Create(&dc).Error
	if err != nil {
		return &CompetitionParticipant{}, err
	}
	return dc, nil
}

func (dc *CompetitionParticipant) FindAllCompetitionParticipants(db *gorm.DB) (*[]CompetitionParticipant, error) {
	ret := []CompetitionParticipant{}
	err := db.Model(&CompetitionParticipant{}).Limit(1000).Find(&ret).Error
	if err != nil {
		return &[]CompetitionParticipant{}, err
	}
	return &ret, err
}

func (dc *CompetitionParticipant) FindCompetitionParticipantByID(db *gorm.DB, uid uint32) (*CompetitionParticipant, error) {
	err := db.Model(CompetitionParticipant{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &CompetitionParticipant{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CompetitionParticipant{}, errors.New("CompetitionParticipant Not Found")
	}
	return dc, err
}

func (dc *CompetitionParticipant) UpdateACompetitionParticipant(db *gorm.DB, uid uint32) (*CompetitionParticipant, error) {
	db = db.Model(&CompetitionParticipant{}).Where("id = ?", uid).Take(&CompetitionParticipant{}).UpdateColumns(
		map[string]interface{}{
			"comment":    dc.Comment,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CompetitionParticipant{}, db.Error
	}
	// This is the display the updated user
	err := db.Model(&CompetitionParticipant{}).Where("id = ?", uid).Take(&dc).Error
	if err != nil {
		return &CompetitionParticipant{}, err
	}
	return dc, nil
}

func (dc *CompetitionParticipant) DeleteACompetitionParticipant(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Model(&CompetitionParticipant{}).Where("id = ?", uid).Take(&CompetitionParticipant{}).Delete(&CompetitionParticipant{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
