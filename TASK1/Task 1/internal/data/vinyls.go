package data

import (
	"fmt"
	"gorm.io/gorm"
	"task1/internal/configs"
	"time"
)

var (
	db *gorm.DB
)

type Vinyl struct {
	gorm.Model
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Year      int32     `json:"year,omitempty"`
}

func init() {
	configs.Connect()
	db := configs.GetDB()
	err := db.AutoMigrate()
	if err != nil {
		fmt.Errorf("badd migration")
	}
}
func CreateVinyl(vinyl *Vinyl) error {
	return db.Create(vinyl).Error
}
func GetAllVinyls() ([]Vinyl, error) {
	var vinyls []Vinyl
	if err := db.Find(&vinyls).Error; err != nil {
		return nil, err
	}
	return vinyls, nil
}
func GetVinylByID(id int64) (*Vinyl, error) {
	var vinyl Vinyl
	if err := db.First(&vinyl, id).Error; err != nil {
		return nil, err
	}
	return &vinyl, nil
}
func DeleteVinyl(id int64) error {
	return db.Delete(&Vinyl{}, id).Error
}
func UpdateVinyl(Id int64, updates any) error {
	return db.Model(&Vinyl{}).Where("id = ?", Id).Updates(updates).Error
}
