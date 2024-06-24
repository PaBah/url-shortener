package dto

type (
	// ShortenRequest - request params for /api/shorten handler
	ShortenRequest struct {
		URL string `json:"url"`
	}

	// ShortenResponse - response params for /api/shorten handler
	ShortenResponse struct {
		Result string `json:"result"`
	}

	// BatchShortenRequest - request params for /api/shorten/batch handler
	BatchShortenRequest struct {
		CorrelationID string `json:"correlation_id"`
		URL           string `json:"original_url"`
	}

	// BatchShortenResponse - response params for /api/shorten/batch handlers
	BatchShortenResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	// UsersURLsResponse - response params for api/user/urls handlers
	UsersURLsResponse struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	// DeleteURLsRequest - request params for shortened URLs deletion handlers
	DeleteURLsRequest []string
)
