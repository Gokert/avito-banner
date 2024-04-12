package annoucement_repo

import (
	"avito-banner/configs"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

//go:generate mockgen -source=banner_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IRepository interface {
}

type Repository struct {
	db *sql.DB
}

func GetPsxRepo(config *configs.DbPsxConfig) (*Repository, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open error: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("sql ping error: %s", err.Error())
	}
	db.SetMaxOpenConns(config.MaxOpenConns)

	return &Repository{db: db}, nil
}
