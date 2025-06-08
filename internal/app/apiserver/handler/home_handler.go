package handler

import (
	"ShorterAPI/internal/domain/shorter"
	"ShorterAPI/internal/domain/shorter/vo"
	"ShorterAPI/internal/utils"
	"ShorterAPI/pkg/dto"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Handler struct {
	Logger *logrus.Logger
	Repo   shorter.Repository
}

func NewHandler(logger *logrus.Logger, Repo shorter.Repository) *Handler {
	return &Handler{
		Logger: logger,
		Repo:   Repo,
	}
}

func (h *Handler) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.UrlRequest

		err := utils.DecodeJSONBody(r, &req)
		if err != nil {
			h.Logger.WithError(err).Error("Error decoding request body")
			respondWithError(w, http.StatusBadRequest, "Bad Request", "invalid request body")
			return
		}

		err = validateUrl(req.LongUrl)
		if err != nil {
			h.Logger.WithError(err).Error("Error validating url")
			respondWithError(w, http.StatusBadRequest, "Bad Request", "invalid url")
			return
		}

		shortUrl := utils.GenerateShortUrl()

		valueObj := vo.NewUrlAliasVO(req.LongUrl, shortUrl)

		err = h.Repo.New(valueObj)
		if err != nil {
			h.Logger.WithError(err).Error("Error creating new obj in table")
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error", "internal server error")
			return
		}
	}
}

func respondWithError(w http.ResponseWriter, status int, errorType string, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(dto.ResponseError{
		ApiError: dto.ApiError{
			Code:    status,
			Message: message,
			Type:    errorType,
		},
	})
	if err != nil {
		http.Error(w, `{"error": "failed to encode error response"}`, http.StatusInternalServerError)
	}
}

func validateUrl(rawUrl string) error {
	if utils.IsEmpty(rawUrl) {
		return errors.New("url is empty")
	}
	parsed, err := url.ParseRequestURI(rawUrl)
	if err != nil || parsed.Host == "" || parsed.Scheme == "" {
		return err
	}
	return nil
}
