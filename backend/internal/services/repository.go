package services

import "gorm.io/gorm"

type Repository interface {
	CreateCalculation(calc Calculation) error
	ReadAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *repo) ReadAllCalculations() ([]Calculation, error) {
	var calculations []Calculation
	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *repo) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation
	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

func (r *repo) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *repo) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
