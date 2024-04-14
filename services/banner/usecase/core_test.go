package usecase

import (
	"avito-banner/pkg/models"
	"avito-banner/services/banner/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserBanner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIRepository(ctrl)

	tagId := uint64(1)
	featureId := uint64(2)
	expectedBanner := &models.UserBanner{Content: "Test Content"}

	mockRepo.EXPECT().GetUserBanner(tagId, featureId).Return(expectedBanner, nil)

	banner, err := mockRepo.GetUserBanner(tagId, featureId)

	assert.NoError(t, err)
	assert.Equal(t, expectedBanner, banner)
}
