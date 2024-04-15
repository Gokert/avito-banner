package banner_repo

import (
	utils "avito-banner/pkg"
	sql_requests "avito-banner/pkg/sql"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserBanner(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql DB: %s", err)
	}
	defer db.Close()

	r := &Repository{db: db}

	tagId := uint64(1)
	featureId := uint64(2)

	expectedContent := "Test Content"

	rows := sqlmock.NewRows([]string{"content"}).AddRow(expectedContent)
	mock.ExpectQuery(sql_requests.GetAllBannersByTagFeature).
		WithArgs(tagId, featureId, utils.DefaultOffset, utils.DefaultLimit).
		WillReturnRows(rows)

	banner, err := r.GetUserBanner(tagId, featureId)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	assert.Equal(t, expectedContent, banner.Content)

	tagId = uint64(999)
	featureId = uint64(999)

	mock.ExpectQuery(sql_requests.GetAllBannersByTagFeature).
		WithArgs(tagId, featureId, utils.DefaultOffset, utils.DefaultLimit).
		WillReturnError(errors.New("database error"))

	_, err = r.GetUserBanner(tagId, featureId)
	assert.True(t, !errors.Is(err, sql.ErrNoRows))

	assert.Nil(t, mock.ExpectationsWereMet())
}
