package dto

type (
	ShortenRequest struct {
		URL string `json:"url"`
	}

	ShortenResponse struct {
		Result string `json:"result"`
	}

	BatchShortenRequest struct {
		CorrelationId string `json:"correlation_id"`
		URL           string `json:"original_url"`
	}

	BatchShortenResponse struct {
		CorrelationId string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}
)
