package tables

const oauth2Table = "oauth2"

const (
	OAuth2Authorization = oauth2Table + "_authorization"
	OAuth2AccessToken   = oauth2Table + "_access_token"
	OAuth2RefreshToken  = oauth2Table + "_refresh_token"
	OAuth2Client        = oauth2Table + "_client"
)

var OAuth2TokenColumns = struct {
	Id        string
	ExpiredAt string
	Data      string
}{
	Id:        "id",
	ExpiredAt: "expired_at",
	Data:      "data",
}

var OAuth2ClientColumns = struct {
	Id   string
	User string
	Data string
}{
	Id:   "id",
	User: "user",
	Data: "data",
}
