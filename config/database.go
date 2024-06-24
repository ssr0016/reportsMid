package config

import (
	"database/sql"
	"fmt"
	"reports/helper"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func ConnectionDB(config *Config) *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.ErrorPanic(err)

	log.Info().Msg("Connected to database!")
	return db
}
