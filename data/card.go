package data

import (
	"github.com/jinzhu/gorm"
	"github.com/caevv/simple-go-prepaid-card/env"
	_ "github.com/jinzhu/gorm/dialects/postgres" // nolint

	"time"
)

type Card struct {
	ID            string  `gorm:"column:id"`
	Amount        float32 `gorm:"column:amount"`
	BlockedAmount float32 `gorm:"column:blocked_amount"`
}

type Amount struct {
	Value float32
}

type Repository struct {
	db *gorm.DB
}

func New() (*Repository, error) {
	// TODO: add a proper await connection with database/sql
	time.Sleep(3 * time.Second)

	db, err := gorm.Open("postgres", env.Settings.DBAddress)
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) TopUp(card Card, amount Amount) (*Card, error) {
	if err := r.db.Find(&card).Error; err != nil {
		return nil, err
	}

	card.Amount += amount.Value

	if err := r.db.Update(&r).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (Card) TableName() string {
	return "cards"
}
