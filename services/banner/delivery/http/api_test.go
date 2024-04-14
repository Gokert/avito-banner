package delivery

import (
	"avito-banner/configs/logger"
	"avito-banner/pkg/models"
	"avito-banner/services/banner/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserBanner(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := logger.GetLogger()

	mockCore := mocks.NewMockICore(ctrl)

	api := &Api{core: mockCore, log: log}

	req, err := http.NewRequest("GET", "/api/v1/user_banner/?tag_id=1&feature_id=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	tagId := uint64(1)
	featureId := uint64(2)
	mockCore.EXPECT().GetUserBanner(tagId, featureId).Return(&models.UserBanner{Content: "Test Content"}, nil)

	api.GetUserBanner(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
