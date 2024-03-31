package schemas

type (
	APIShortenRequestSchema struct {
		URL string `json:"url"`
	}

	APIShortenResponseSchema struct {
		Result string `json:"result"`
	}
)
