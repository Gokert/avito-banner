package delivery

import (
	utils "avito-banner/pkg"
	"avito-banner/pkg/middleware"
	"avito-banner/pkg/models"
	httpResponse "avito-banner/pkg/response"
	"avito-banner/services/banner/usecase"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"time"
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

	api.mx.Handle("/api/v1/user_banner", middleware.AuthCheck(middleware.MethodCheck(http.HandlerFunc(api.GetUserBanner), http.MethodGet, log), core, log))
	api.mx.Handle("/api/v1/banner", middleware.AuthCheck(middleware.GetRole(http.HandlerFunc(api.GetOrCreateBanner), core, log), core, log))
	api.mx.Handle("/api/v1/banner/", middleware.AuthCheck(middleware.CheckRole(http.HandlerFunc(api.EditOrDeleteBanner), core, log), core, log))

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
	response := &models.Response{Status: http.StatusOK, Body: nil}
	timeNow := time.Now()

	tagId, err := strconv.ParseUint(r.URL.Query().Get("tag_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Body = models.Error{Message: "no have tag_id query param"}
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	featureId, err := strconv.ParseUint(r.URL.Query().Get("feature_id"), 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Body = models.Error{Message: "no have feature_id query param"}
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	lastVersion, err := strconv.ParseBool(r.URL.Query().Get("last_version"))
	if err != nil {
		lastVersion = false
	}

	banner, find, err := a.core.GetUserBanner(r.Context(), tagId, featureId, lastVersion)
	if err != nil {
		a.log.Errorf("Get user banner error: %s", err.Error())
		response.Status = http.StatusInternalServerError
		response.Body = models.Error{Message: "server error"}
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	if !find {
		response.Status = http.StatusNotFound
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	response.Body = banner

	httpResponse.SendResponse(w, r, response, timeNow, a.log)
}

func (a *Api) GetOrCreateBanner(w http.ResponseWriter, r *http.Request) {
	response := &models.Response{Status: http.StatusOK, Body: nil}
	timeNow := time.Now()

	if http.MethodGet == r.Method {
		tagIdStr := r.URL.Query().Get("tag_id")
		if tagIdStr == "" {
			tagIdStr = "0"
		}
		tagId, err := strconv.ParseUint(tagIdStr, 10, 64)
		if err != nil {
			response.Status = http.StatusBadRequest
			response.Body = models.Error{Message: "parse tag id error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		featureIdStr := r.URL.Query().Get("feature_id")
		if featureIdStr == "" {
			featureIdStr = "0"
		}
		featureId, err := strconv.ParseUint(featureIdStr, 10, 64)
		if err != nil {
			response.Status = http.StatusBadRequest
			response.Body = models.Error{Message: "parse feature id error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		offset, err := strconv.ParseUint(r.URL.Query().Get("offset"), 10, 64)
		if err != nil {
			offset = utils.DefaultOffset
		}

		limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 64)
		if err != nil {
			limit = utils.DefaultLimit
		}

		getAllBanner := false
		if r.Context().Value(middleware.UserRoleKey) == "admin" {
			getAllBanner = true
		}

		banners, err := a.core.GetBanners(r.Context(), tagId, featureId, getAllBanner, offset, limit)
		if err != nil {
			a.log.Errorf("Get banners error: %s", err.Error())
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "server error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		response.Body = banners
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	if http.MethodPost == r.Method {
		if r.Context().Value(middleware.UserRoleKey) != "admin" {
			response.Status = http.StatusForbidden
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		var banner models.BannerRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Status = http.StatusBadRequest
			response.Body = models.Error{Message: "read body error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		err = json.Unmarshal(body, &banner)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "json unmarshal error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		err = a.core.CreateBanner(r.Context(), &banner)
		if err != nil {
			a.log.Errorf("Create banner error: %s", err.Error())
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "server error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		response.Status = http.StatusCreated

		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	response.Status = http.StatusMethodNotAllowed
	httpResponse.SendResponse(w, r, response, timeNow, a.log)
}

func (a *Api) EditOrDeleteBanner(w http.ResponseWriter, r *http.Request) {
	response := &models.Response{Status: http.StatusOK, Body: nil}
	timeNow := time.Now()

	bannerId, err := strconv.ParseUint(r.URL.Path[len("/api/v1/banner/"):], 10, 64)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Body = models.Error{Message: "id query param error"}
		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	if http.MethodPatch == r.Method {
		var banner models.BannerRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Status = http.StatusBadRequest
			response.Body = models.Error{Message: "read body error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		err = json.Unmarshal(body, &banner)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "json unmarshal error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		banner.BannerId = bannerId
		res, err := a.core.UpdateBanner(r.Context(), &banner)
		if err != nil {
			a.log.Errorf("Update banner error: %s", err.Error())
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "server error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		if !res {
			response.Status = http.StatusNotFound
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	if http.MethodDelete == r.Method {
		res, err := a.core.DeleteBanner(r.Context(), bannerId)
		if err != nil {
			a.log.Errorf("Delete banner error: %s", err.Error())
			response.Status = http.StatusInternalServerError
			response.Body = models.Error{Message: "server error"}
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		if !res {
			response.Status = http.StatusNotFound
			httpResponse.SendResponse(w, r, response, timeNow, a.log)
			return
		}

		httpResponse.SendResponse(w, r, response, timeNow, a.log)
		return
	}

	response.Status = http.StatusMethodNotAllowed
	httpResponse.SendResponse(w, r, response, timeNow, a.log)
}
