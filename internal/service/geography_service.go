package service

import (
	"time"

	"github.com/Hitesh3602/master_geography/internal/db"
	"github.com/Hitesh3602/master_geography/internal/model"
)

type GeographyService interface {
	CreateGeography(geography *model.Geography) error
	GetGeographies() ([]*model.Geography, error)
	GetGeographyByID(id int64) (*model.Geography, error)
	UpdateGeography(geography *model.Geography) error
	DeleteGeography(id int64) error
}

type geographyService struct {
	repo db.GeographyRepository
}

func NewGeographyService(repo db.GeographyRepository) GeographyService {
	return &geographyService{repo: repo}
}

func (s *geographyService) CreateGeography(geography *model.Geography) error {
	geography.CreatedAt = time.Now()
	geography.UpdatedAt = time.Now()
	return s.repo.Create(geography)
}

func (s *geographyService) GetGeographies() ([]*model.Geography, error) {
	return s.repo.GetAll()
}

func (s *geographyService) GetGeographyByID(id int64) (*model.Geography, error) {
	return s.repo.GetByID(id)
}

func (s *geographyService) UpdateGeography(geography *model.Geography) error {
	return s.repo.Update(geography)
}

func (s *geographyService) DeleteGeography(id int64) error {
	return s.repo.Delete(id)
}
