package models

import (
    "gorm.io/gorm"
)

type Serie struct {
    gorm.Model
    Title            string `json:"title" gorm:"type:varchar(100);not null"`
    Status           string `json:"status" gorm:"type:varchar(50);not null"`
    LastEpisodeWatched int `json:"lastEpisodeWatched" gorm:"default:0"`
    TotalEpisodes    int    `json:"totalEpisodes" gorm:"default:0"`
    Ranking          int    `json:"ranking" gorm:"default:0"`
}

func GetAllSeries(db *gorm.DB) ([]Serie, error) {
    var series []Serie
    result := db.Find(&series)
    return series, result.Error
}

func GetSeriesByID(db *gorm.DB, id int) (*Serie, error) {
    var serie Serie
    result := db.First(&serie, id)
    return &serie, result.Error
}

func CreateSerie(db *gorm.DB, serie *Serie) error {
    result := db.Create(serie)
    return result.Error
}

func UpdateSerie(db *gorm.DB, id int, serie *Serie) error {
    var existing Serie
    db.First(&existing, id)
    result := db.Model(&existing).Updates(serie)
    return result.Error
}

func DeleteSerie(db *gorm.DB, id int) error {
    result := db.Delete(&Serie{}, id)
    return result.Error
}