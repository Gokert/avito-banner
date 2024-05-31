package banner_repo

import (
	"avito-banner/configs"
	utils "avito-banner/pkg"
	"avito-banner/pkg/models"
	sql_requests "avito-banner/pkg/sql"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

//go:generate mockgen -source=banner_repo.go -destination=../../mocks/repo_mock.go -package=mocks
type IRepository interface {
	GetUserBanner(ctx context.Context, tagId uint64, featureId uint64, lastVersion bool) (*models.UserBanner, error)
	GetBanners(ctx context.Context, tagId uint64, featureId uint64, getAllBanners bool, offset uint64, limit uint64) (*[]models.BannerResponse, error)
	CheckFeature(ctx context.Context, featureId uint64) (bool, error)
	CreateBanner(ctx context.Context, banner *models.BannerRequest) error
	CheckBanner(ctx context.Context, bannerId uint64) (bool, error)
	DeleteBanner(ctx context.Context, bannerId uint64) (bool, error)
	UpdateBanner(ctx context.Context, banner *models.BannerRequest) (bool, error)
}

type Repository struct {
	db *sql.DB
}

func GetPsxRepo(config *configs.DbPsxConfig, logger *logrus.Logger) (*Repository, error) {
	dsn := fmt.Sprintf("user=%s dbname=%s password= %s host=%s port=%d sslmode=%s",
		config.User, config.Dbname, config.Password, config.Host, config.Port, config.Sslmode)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql open error: %s", err.Error())
	}

	repo := &Repository{db: db}

	errs := make(chan error)
	go func() {
		errs <- repo.pingDb(3, logger)
	}()

	if err := <-errs; err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	logger.Info("Successfully connected to database")

	return repo, nil
}

func (r *Repository) pingDb(timer uint32, log *logrus.Logger) error {
	var err error
	var retries int

	for retries < utils.MaxRetries {
		err = r.db.Ping()
		if err == nil {
			return nil
		}

		retries++
		log.Errorf("sql ping error: %s", err.Error())
		time.Sleep(time.Duration(timer) * time.Second)
	}

	return fmt.Errorf("sql max pinging error: %s", err)
}

func (r *Repository) GetUserBanner(ctx context.Context, tagId uint64, featureId uint64, lastVersion bool) (*models.UserBanner, error) {
	var banner models.UserBanner
	var sqlString string

	if lastVersion {
		sqlString = sql_requests.GetUserBannerLast
	} else {
		sqlString = sql_requests.GetUserBanner
	}

	err := r.db.QueryRowContext(ctx, sqlString, tagId, featureId).Scan(&banner.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("select user banner error: %s", err)
	}

	return &banner, nil
}

func (r *Repository) GetTagsIdOfBanner(ctx context.Context, bannerId uint64) (*[]uint64, error) {
	var tags []uint64

	rows, err := r.db.QueryContext(ctx, "SELECT id FROM tags JOIN banner_tags ON tags.id = banner_tags.id_tag WHERE banner_tags.id_banner = $1", bannerId)
	if err != nil {
		return nil, fmt.Errorf("select tags error: %s", err)
	}

	for rows.Next() {
		var tag uint64

		err := rows.Scan(&tag)
		if err != nil {
			return nil, fmt.Errorf("select tags error: %s", err)
		}

		tags = append(tags, tag)
	}

	return &tags, nil
}

func (r *Repository) CheckBanner(ctx context.Context, bannerId uint64) (bool, error) {
	var id uint64

	err := r.db.QueryRowContext(ctx, "SELECT id FROM banners WHERE id = $1", bannerId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("select id banner error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) CheckFeature(ctx context.Context, featureId uint64) (bool, error) {
	var id uint64

	err := r.db.QueryRowContext(ctx, "SELECT id FROM features WHERE id = $1", featureId).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, fmt.Errorf("select id feature error: %s", err.Error())
	}

	return true, nil
}

func (r *Repository) UpdateBanner(ctx context.Context, banner *models.BannerRequest) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	_, err = tx.ExecContext(ctx, "UPDATE banners SET id_feature = $1 WHERE id = $2", banner.FeatureId, banner.BannerId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, fmt.Errorf("rollback error: %s", err.Error())
		}
		return false, fmt.Errorf("update banner error: %s", err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO versions(id_banner, is_active, content) VALUES ($1, $2, $3);",
		banner.BannerId, banner.IsActive, banner.Content)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, fmt.Errorf("rollback error: %s", err.Error())
		}
		return false, fmt.Errorf("update banner error: %s", err)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM banner_tags WHERE id_banner = $1 AND id_tag = ANY($2)", banner.BannerId, pq.Array(banner.TagIds))
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, fmt.Errorf("rollback error: %s", err.Error())
		}
		return false, fmt.Errorf("delete old tags error: %s", err)
	}

	for _, tagId := range banner.TagIds {
		_, err = tx.ExecContext(ctx, "INSERT INTO banner_tags(id_tag, id_banner) VALUES ($1, $2) ON CONFLICT DO NOTHING", tagId, banner.BannerId)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false, fmt.Errorf("rollback error: %s", err.Error())
			}
			return false, fmt.Errorf("insert tag for banner error: %s", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return false, fmt.Errorf("rollback error: %s", err.Error())
		}
		return false, fmt.Errorf("failed to commit transaction: %s", err)
	}

	return true, nil
}

func (r *Repository) GetBanners(ctx context.Context, tagId uint64, featureId uint64, getAllBanners bool, offset uint64, limit uint64) (*[]models.BannerResponse, error) {
	var banners []models.BannerResponse
	var rows *sql.Rows
	var err error

	query := ""
	args := []interface{}{offset, limit}

	switch {
	case tagId != 0 && featureId != 0:
		if getAllBanners {
			query = sql_requests.GetBannersByTagFeature
		} else {
			query = sql_requests.GetBannersByTag
		}
		args = append([]interface{}{tagId, featureId}, args...)
	case tagId != 0:
		if getAllBanners {
			query = sql_requests.GetAllBannersByTag
		} else {
			query = sql_requests.GetBannersByTag
		}
		args = append([]interface{}{tagId}, args...)
	case featureId != 0:
		if getAllBanners {
			query = sql_requests.GetAllBannersFeature
		} else {
			query = sql_requests.GetBannersFeature
		}
		args = append([]interface{}{featureId}, args...)
	default:
		if getAllBanners {
			query = sql_requests.GetAllBanners
		} else {
			query = sql_requests.GetBanners
		}
	}

	rows, err = r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select get banners error: %s", err.Error())
	}

	for rows.Next() {
		var banner models.BannerResponse

		err := rows.Scan(&banner.BannerId, &banner.FeatureId, &banner.Content, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan get banners error: %s", err.Error())
		}

		tags, err := r.GetTagsIdOfBanner(ctx, banner.BannerId)
		if err != nil {
			return nil, fmt.Errorf("get tags error: %s", err.Error())
		}

		banner.TagIds = *tags
		banners = append(banners, banner)
	}

	return &banners, nil
}

func (r *Repository) CreateBanner(ctx context.Context, banner *models.BannerRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	var bannerId uint64
	err = tx.QueryRowContext(ctx, "INSERT INTO banners(id_feature) VALUES ($1) RETURNING id", banner.FeatureId).Scan(&bannerId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return fmt.Errorf("rollback error: %s", err.Error())
		}
		return fmt.Errorf("insert banner error: %s", err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO versions(id_banner, content) VALUES ($1, $2)", bannerId, banner.Content)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return fmt.Errorf("rollback error: %s", err.Error())
		}
		return fmt.Errorf("insert version error: %s", err)
	}

	for _, tagId := range banner.TagIds {
		_, err = tx.ExecContext(ctx, "INSERT INTO banner_tags(id_tag, id_banner) VALUES ($1, $2)", tagId, bannerId)
		if err != nil {
			err := tx.Rollback()
			if err != nil && !errors.Is(err, sql.ErrTxDone) {
				return fmt.Errorf("insert banner tags error: %s", err.Error())
			}
			return fmt.Errorf("insert tag for banner error: %s", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %s", err.Error())
	}

	return nil
}

func (r *Repository) DeleteBanner(ctx context.Context, bannerId uint64) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, fmt.Errorf("failed to begin transaction: %s", err.Error())
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM banners WHERE id = $1", bannerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("delete banner error: %s", err.Error())
	}

	if err := tx.Commit(); err != nil {
		return false, fmt.Errorf("failed to commit transaction: %s", err.Error())
	}

	return true, nil
}
