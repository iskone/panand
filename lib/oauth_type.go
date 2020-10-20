package lib
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
func (r OauthErr) HasErr() error {
	if r.Error=="" {
		return nil
	}
	return Err{
		Code: r.Error,
		Desc: r.ErrorDescription,
	}
}
