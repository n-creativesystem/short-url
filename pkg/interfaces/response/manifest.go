package response

type WebUIManifest struct {
	CsrfTokenBase bool   `json:"token_base"`
	HeaderName    string `json:"header_name"`
}
