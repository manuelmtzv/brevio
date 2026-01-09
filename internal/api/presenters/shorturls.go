package presenters

import (
	"github.com/manuelmtzv/brevio/internal/models"
)

type ShortURLResponse struct {
	URL string `json:"url"`
}

func ShortURL(u *models.ShortURL, urls URLBuilder) ShortURLResponse {
	return ShortURLResponse{
		URL: urls.ShortURL(u.Code),
	}
}
