package banner_repo

import (
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
	mock.ExpectQuery("SELECT content FROM banners JOIN banner_tags on banners.id = banner_tags.id_banner WHERE banner_tags.id_tag = \\$1 AND banners.id_feature = \\$2").
		WithArgs(tagId, featureId).
		WillReturnRows(rows)

	banner, err := r.GetUserBanner(tagId, featureId)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	assert.Equal(t, expectedContent, banner.Content)

	tagId = uint64(999)
	featureId = uint64(999)

	mock.ExpectQuery("SELECT content FROM banners JOIN banner_tags on banners.id = banner_tags.id_banner WHERE banner_tags.id_tag = \\$1 AND banners.id_feature = \\$2").
		WithArgs(tagId, featureId).
		WillReturnError(errors.New("database error"))

	_, err = r.GetUserBanner(tagId, featureId)
	assert.True(t, true)

	assert.Nil(t, mock.ExpectationsWereMet())
}
