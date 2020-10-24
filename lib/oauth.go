package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Auth interface {
	Login(state string) string
	AccessToken(code string) (AccessToken, error)
	UseWap(bool)
	RefreshToken(string) (RefreshToken, error)
	Name() string
}
type client struct {
	AppName     string
	ClientId    string
	AppKey      string
	RedirectUri string
	useWap      bool
}

func (c client) Login(state string) string {
	var api = OauthDefaultHost
	if c.useWap {
		api = OauthWapHost
	}
	return fmt.Sprintf("%s?response_type=%s&client_id=%s&redirect_uri=%s&scope=%s&state=%s",
		api,
		"code",
		c.ClientId,
		c.RedirectUri,
		OauthScope,
		state)
}

func (c client) AccessToken(code string) (AccessToken, error) {
	client := &http.Client{}
	v := url.Values{}
	v.Add("grant_type", "authorization_code")
	v.Add("code", code)
	v.Add("redirect_uri", c.RedirectUri)
	req, _ := http.NewRequest(http.MethodPost, PanApi+GetTokenApi, strings.NewReader(v.Encode()))
	req.SetBasicAuth(c.ClientId, c.AppKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, e := client.Do(req)
	var at = AccessToken{}
	if e != nil {
		return at, e
	}

	if e = json.NewDecoder(res.Body).Decode(&at); e != nil {
		return at, e
	}
	return at, nil
}

func (c client) RefreshToken(rtoken string) (RefreshToken, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, PanApi+RefreshTokenApi, strings.NewReader("grant_type=refresh_token&refresh_token="+rtoken))
	req.SetBasicAuth(c.ClientId, c.AppKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, e := client.Do(req)
	var rt = RefreshToken{}
	if e != nil {
		return rt, e
	}
	if e = json.NewDecoder(res.Body).Decode(&rt); e != nil {
		return rt, e
	}
	if e = rt.HasErr(); e != nil {
		return rt, e
	}
	return rt, nil
}
func (c *client) UseWap(b bool) {
	c.useWap = b
}
func (c *client) Name() string {
	return c.AppName
}
func NewClientOauth(name string, id string, key string, redirectUri string) Auth {
	return &client{
		AppName:     name,
		ClientId:    id,
		AppKey:      key,
		RedirectUri: redirectUri,
	}
}
