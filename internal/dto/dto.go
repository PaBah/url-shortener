package dto

type (
	ShortenRequest struct {
		URL string `json:"url"`
	}

	ShortenResponse struct {
		Result string `json:"result"`
	}

	BatchShortenRequest struct {
		CorrelationID string `json:"correlation_id"`
		URL           string `json:"original_url"`
	}

	BatchShortenResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	UsersURLsResponse struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	DeleteURLsRequest []string
)
