package lib

import "encoding/xml"

type Code struct {
	CODE  string `json:"CODE"`
	State string `json:"state"`
	OauthErr
}
type OauthErr struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
type AccessToken struct {
	RefreshToken string `json:"refresh_token"`
	Token
	OauthErr
}
type RefreshToken struct {
	Token
	OauthErr
}
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
type IDs struct {
	ID     []string `xml:"ID"`
}
type Result struct {
	XMLName     xml.Name `xml:"result"`
	Text        string   `xml:",chardata"`
	ResultCode  string   `xml:"resultCode,attr"`
	Desc        string   `xml:"desc,attr"`
	RetCode     string   `xml:"retCode"`
}
