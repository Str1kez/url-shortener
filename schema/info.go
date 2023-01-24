package schema

type InfoRequest struct {
	SecretKey string `uri:"secret_key" binding:"required,uuid"`
}

type InfoResponse struct {
	ShortUrl       string `json:"short_url"`
	LongUrl        string `json:"long_url"`
	NumberOfClicks int    `json:"number_of_clicks"`
	CreatedAt      string `json:"dt_created"`
}
