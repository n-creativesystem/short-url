package response

type GenerateShortURL struct {
	URL string `json:"url"`
}

type Shorts struct {
	Key       string `json:"key"`
	URL       string `json:"url"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
