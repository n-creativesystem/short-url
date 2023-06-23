package oauth2client

import "github.com/n-creativesystem/short-url/pkg/utils/credentials"

type Client struct {
	ID      string
	Secret  credentials.EncryptString
	Domain  credentials.EncryptString
	Public  bool
	UserID  string
	AppName string
}

func NewClient(id, secret, domain string, public bool, userID, appName string) Client {
	return Client{
		ID:      id,
		Secret:  credentials.NewEncryptString(secret),
		Domain:  credentials.NewEncryptString(domain),
		Public:  public,
		UserID:  userID,
		AppName: appName,
	}
}

func (c *Client) GetAppName() string {
	return c.AppName
}

func (c *Client) GetEncryptSecret() credentials.EncryptString {
	return c.Secret
}

func (c *Client) GetEncryptDomain() credentials.EncryptString {
	return c.Domain
}

// GetID client id
func (c *Client) GetID() string {
	return c.ID
}

// GetSecret client secret
func (c *Client) GetSecret() string {
	return c.Secret.UnmaskedString()
}

// GetDomain client domain
func (c *Client) GetDomain() string {
	return c.Domain.UnmaskedString()
}

// IsPublic public
func (c *Client) IsPublic() bool {
	return c.Public
}

// GetUserID user id
func (c *Client) GetUserID() string {
	return c.UserID
}
