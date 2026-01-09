package handlers

import (
	"net/http"

	appErrors "github.com/manuelmtzv/brevio/internal/errors"

	"github.com/manuelmtzv/brevio/internal/api/presenters"
	"github.com/manuelmtzv/brevio/internal/api/services"
	"github.com/manuelmtzv/brevio/internal/http/render"
	"github.com/manuelmtzv/brevio/internal/http/request"
	"github.com/manuelmtzv/brevio/internal/http/validate"
	"github.com/manuelmtzv/brevio/internal/i18n"
	"go.uber.org/zap"
)

type ShortURLHandlerDeps struct {
	Service    services.ShortURLService
	Localizer  i18n.Localizer
	URLBuilder presenters.URLBuilder
	Logger     *zap.SugaredLogger
}

type CreateShortURLRequest struct {
	Target string `json:"target" validate:"required,url"`
}

type ShortURLHandler struct {
	service    services.ShortURLService
	localizer  i18n.Localizer
	urlBuilder presenters.URLBuilder
	logger     *zap.SugaredLogger
}

func NewShortURLHandler(deps ShortURLHandlerDeps) *ShortURLHandler {
	return &ShortURLHandler{
		service:    deps.Service,
		localizer:  deps.Localizer,
		urlBuilder: deps.URLBuilder,
		logger:     deps.Logger,
	}
}

func (h *ShortURLHandler) HandleCreate(w http.ResponseWriter, r *http.Request) error {
	var payload CreateShortURLRequest
	if err := request.DecodeJSON(w, r, &payload); err != nil {
		return err
	}

	if err := validate.Struct(payload); err != nil {
		return appErrors.Validation(err)
	}

	shortURL, err := h.service.Create(r.Context(), services.CreateShortURLInput{
		Target: payload.Target,
	})
	if err != nil {
		return err
	}

	return render.JSON(w, http.StatusCreated, presenters.ShortURL(shortURL, h.urlBuilder))
}
