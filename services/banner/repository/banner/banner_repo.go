package banner_repo

import (
	"avito-banner/configs"
	"avito-banner/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

//go:generate mockgen -source=banner_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IRepository interface {
	GetUserBanner(tagId uint64, featureId uint64) (*models.UserBanner, error)
	GetBanners(tagId uint64, featureId uint64, offset uint64, limit uint64) (*[]models.BannerResponse, error)
	CreateBanner(banner *models.BannerRequest) error
	CheckBanner(bannerId uint64) (bool, error)
	DeleteBanner(bannerId uint64) (bool, error)
	UpdateBanner(banner *models.BannerRequest) (bool, error)
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

func (r *Repository) CheckBanner(bannerId uint64) (bool, error) {
	var id uint64

	err := r.db.QueryRow("SELECT id FROM banners WHERE id = $1", bannerId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("select id banner error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) UpdateBanner(banner *models.BannerRequest) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	_, err = tx.Exec("UPDATE banners SET id_feature = $1, title = $2, text = $3, url = $4, is_active = $5, updated_at = now() WHERE id = $6",
		banner.FeatureId, banner.Content["title"], banner.Content["text"], banner.Content["url"], banner.IsActive, banner.BannerId)
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("update banner error: %s", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("failed to commit transaction: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) GetUserBanner(tagId uint64, featureId uint64) (*models.UserBanner, error) {
	var banner models.UserBanner

	err := r.db.QueryRow("SELECT banners.title, banners.text, banners.url FROM banners JOIN banner_tags on banners.id = banner_tags.id_banner WHERE banner_tags.id_tag = $1 AND banners.id_feature = $2", tagId, featureId).Scan(&banner.Title, &banner.Text, &banner.Url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("select user banner error: %s", err)
	}

	return &banner, nil
}

func (r *Repository) GetBanners(tagId uint64, featureId uint64, offset uint64, limit uint64) (*[]models.BannerResponse, error) {
	var banners []models.BannerResponse
	var rows *sql.Rows
	var err error

	if tagId != 0 && featureId != 0 {
		rows, err = r.db.Query("SELECT id, id_feature, is_active, created_at, updated_at FROM banners JOIN banner_tags ON banner_tags.id_banner = banners.id WHERE banner_tags.id_tag = $1 AND banners.id_feature = $2 OFFSET $3 LIMIT $4", tagId, featureId, offset, limit)
	} else if tagId != 0 {
		rows, err = r.db.Query("SELECT id, id_feature, is_active, created_at, updated_at FROM banners JOIN banner_tags ON banner_tags.id_banner = banners.id WHERE banner_tags.id_tag = $1 OFFSET $2 LIMIT $3", tagId, offset, limit)
	} else if featureId != 0 {
		rows, err = r.db.Query("SELECT id, id_feature, is_active, created_at, updated_at FROM banners WHERE banners.id_feature = $1 OFFSET $2 LIMIT $3", featureId, offset, limit)
	} else {
		rows, err = r.db.Query("SELECT id, id_feature, is_active, created_at, updated_at FROM banners OFFSET $1 LIMIT $2", offset, limit)
	}

	if err != nil {
		return nil, fmt.Errorf("select get banners error: %s", err.Error())
	}

	for rows.Next() {
		var banner models.BannerResponse

		err := rows.Scan(&banner.BannerId, &banner.FeatureId, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan get banners error: %s", err.Error())
		}

		banners = append(banners, banner)
	}

	return &banners, nil
}

func (r *Repository) CreateBanner(banner *models.BannerRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	err = tx.QueryRow("INSERT INTO banners(id_feature, is_active, title, text, url) VALUES ($1, $2, $3, $4, $5) RETURNING id", banner.FeatureId, banner.IsActive, banner.Content["title"], banner.Content["text"], banner.Content["url"]).Scan(&banner.BannerId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("insert into banner error: %s", err.Error())
	}

	for _, tagId := range banner.TagIds {
		_, err = tx.Exec("INSERT INTO banner_tags(id_tag, id_banner) VALUES ($1, $2)", tagId, banner.BannerId)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("insert into banner_ error: %s", err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %s", err.Error())
	}

	return nil
}

func (r *Repository) DeleteBanner(bannerId uint64) (bool, error) {
	_, err := r.db.Exec("DELETE FROM banners WHERE id = $1", bannerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("delete banner error: %s", err.Error())
	}

	return true, nil
}
