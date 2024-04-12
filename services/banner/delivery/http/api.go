package delivery

import (
	"avito-banner/pkg/middleware"
	"avito-banner/pkg/models"
	httpResponse "avito-banner/pkg/response"
	"avito-banner/services/banner/usecase"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Api struct {
	log  *logrus.Logger
	mx   *http.ServeMux
	core usecase.ICore
}

func GetApi(core *usecase.Core, log *logrus.Logger) *Api {
	api := &Api{
		core: core,
		log:  log,
		mx:   http.NewServeMux(),
	}

	api.mx.Handle("/user_banner", middleware.AuthCheck(middleware.MethodCheck(http.HandlerFunc(api.GetUserBanner), http.MethodGet, log), core, log))
	api.mx.Handle("/banner", middleware.AuthCheck(http.HandlerFunc(api.GetBanners), core, log))
	api.mx.Handle("/banner/{id}", middleware.AuthCheck(http.HandlerFunc(api.EditOrDeleteBanner), core, log))

	return api
}

func (a *Api) ListenAndServe(port string) error {
	err := http.ListenAndServe(":"+port, a.mx)
	if err != nil {
		a.log.Errorf("listen error: %s", err.Error())
		return err
	}

	return nil
}

func (a *Api) GetUserBanner(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: http.StatusOK, Body: nil}

	tagId, err := strconv.ParseUint(r.URL.Query().Get("tag_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		httpResponse.SendResponse(w, r, &response, a.log)
	}

	featureId, err := strconv.ParseUint(r.URL.Query().Get("feature_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		httpResponse.SendResponse(w, r, &response, a.log)
	}

	fmt.Println(tagId, featureId)

	httpResponse.SendResponse(w, r, &response, a.log)
}

func (a *Api) GetBanners(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: http.StatusOK, Body: nil}

	httpResponse.SendResponse(w, r, &response, a.log)
}

func (a *Api) EditOrDeleteBanner(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: http.StatusOK, Body: nil}

	httpResponse.SendResponse(w, r, &response, a.log)
}
