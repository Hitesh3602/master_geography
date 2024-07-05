package db

import "github.com/Hitesh3602/master_geography/internal/model"

type GeographyRepository interface {
	Create(geo *model.Geography) error
	GetAll() ([]*model.Geography, error)
	GetByID(id int64) (*model.Geography, error)
	Update(geo *model.Geography) error
	Delete(id int64) error
}
