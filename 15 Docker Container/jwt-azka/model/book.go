package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `json:"title" form:"title"`
	AuthorID    uint           `json:"author_id" form:"author_id"`
	PublisherID uint           `json:"publisher_id" form:"publisher_id"`
	Price       int            `json:"price" form:"price"`
	Page        int            `json:"page" form:"page"`
	Weight      int            `json:"weight" form:"weight"`
	PublishedAt datatypes.Date `json:"published_at" form:"published_at"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Categorie struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CategoryBook struct {
	CategoryID uint `gorm:"primaryKey;autoIncrement:false"`
	BookID     uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Publisher struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" form:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
