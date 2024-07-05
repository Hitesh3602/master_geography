package db

import (
	"database/sql"

	"github.com/Hitesh3602/master_geography/internal/model"
)

type PostgresGeographyRepository struct {
	DB *sql.DB
}

func NewPostgresGeographyRepository(db *sql.DB) *PostgresGeographyRepository {
	return &PostgresGeographyRepository{DB: db}
}

func (r *PostgresGeographyRepository) Create(geo *model.Geography) error {
	query := `INSERT INTO master_geography (type, name, value, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.DB.QueryRow(query, geo.Type, geo.Name, geo.Value, geo.Metadata, geo.CreatedAt, geo.UpdatedAt).Scan(&geo.ID)
}

func (r *PostgresGeographyRepository) GetAll() ([]*model.Geography, error) {
	query := `SELECT id, type, name, value, metadata, created_at, updated_at FROM master_geography ORDER BY id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	geographies := []*model.Geography{}
	for rows.Next() {
		geo := new(model.Geography)
		err := rows.Scan(&geo.ID, &geo.Type, &geo.Name, &geo.Value, &geo.Metadata, &geo.CreatedAt, &geo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		geographies = append(geographies, geo)
	}
	return geographies, nil
}

func (r *PostgresGeographyRepository) GetByID(id int64) (*model.Geography, error) {
	query := `SELECT id, type, name, value, metadata, created_at, updated_at FROM master_geography WHERE id = $1`
	geo := new(model.Geography)
	err := r.DB.QueryRow(query, id).Scan(&geo.ID, &geo.Type, &geo.Name, &geo.Value, &geo.Metadata, &geo.CreatedAt, &geo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return geo, nil
}

func (r *PostgresGeographyRepository) Update(geo *model.Geography) error {
	query := `UPDATE master_geography SET type = $1, name = $2, value = $3, metadata = $4, updated_at = $5 WHERE id = $6`
	_, err := r.DB.Exec(query, geo.Type, geo.Name, geo.Value, geo.Metadata, geo.UpdatedAt, geo.ID)
	return err
}

func (r *PostgresGeographyRepository) Delete(id int64) error {
	query := `DELETE FROM master_geography WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
