package lib

import (
	"encoding/xml"
)

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
	RToken string `json:"refresh_token"`
	RefreshToken
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
type Err struct {
	Code string
	Desc string
}
func (e Err) Error() string {
	msg:= "code: "+e.Code
	if e.Desc!="" {
		msg += ",desc: "+e.Desc

	}
	return msg
}

func (r Result) HasErr() error {
	e:=Err{
		Desc: r.Desc,
	}
	if r.RetCode!="0" &&r.RetCode!="" {
		e.Code = r.RetCode
		return e
	}
	if r.ResultCode !="0" &&r.ResultCode!="" {
		e.Code = r.ResultCode
		return e
	}
	return nil
}
func (r OauthErr) HasErr() error {
	if r.Error=="" {
		return nil
	}
	return Err{
		Code: r.Error,
		Desc: r.ErrorDescription,
	}
}
