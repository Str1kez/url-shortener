package schema

type InfoResponse struct {
	ShortUrl       string `json:"short_url"`
	LongUrl        string `json:"long_url"`
	NumberOfClicks int    `json:"number_of_clicks"`
	CreatedAt      string `json:"dt_created"`
}
