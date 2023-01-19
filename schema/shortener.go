package schema

type ShortenerRequest struct {
	Url string `json:"url" binding:"required,url"`
}

type ShortenerResponse struct {
	ShortUrl  string `json:"short_url"`
	SecretKey string `json:"secret_key"`
}
