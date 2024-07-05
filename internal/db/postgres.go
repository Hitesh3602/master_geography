package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
    "github.com/Hitesh3602/master_geography/internal/config"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    // Check if the database is connected
    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Database connected successfully")
    return db, nil
}
