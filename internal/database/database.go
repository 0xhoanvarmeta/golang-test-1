package database

import (
	"database/sql"
	"fmt"
	"os"
	"test-1/internal/configs"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
)

// Global Database connection
var Db *sql.DB

func buildDatabaseUrl(config *configs.Config) string {
	//return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
	//	config.Database.Username,
	//	config.Database.Password,
	//	config.Database.Host,
	//	config.Database.Port,
	//	config.Database.DBName,
	//)

	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.DBName,
		config.Database.Password,
	)
}

func CreateDatabaseConnection(config *configs.Config) {
	//dbPool, err := pgxpool.New(context.Background(), buildDatabaseUrl(config))
	//if err != nil {
	//	log.Error().Err(err).Msg(fmt.Sprintf("Unable to connect to database: %v", err))
	//	os.Exit(1)
	//}
	//
	//log.Info().Msg("Connected to database")
	//
	//defer dbPool.Close()

	db, errSql := sql.Open("pgx", buildDatabaseUrl(config))
	if errSql != nil {
		log.Error().Err(errSql).Msg(fmt.Sprintf("Unable to connect to database: %v", errSql))
		os.Exit(1)
	}

	Db = db
	log.Info().Msg("Successfully connected to database!")

}

func GetContextWithTimeout() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	return ctx, cancel
}
