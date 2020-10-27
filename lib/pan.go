package lib

import (
	"fmt"
	"net/http"
	"time"
)

type Panand struct {
	Auth Auth
	AT *AccessToken
}

func (p *Panand) addHeader(r *http.Request) error {
	if p.AT.ExpiresIn < int(time.Now().Unix()) {
		if rt, e := p.Auth.RefreshToken(p.AT.RToken); e != nil {
			return e
		} else {
			p.AT.Token = rt.Token
		}
	}
	fmt.Println(p.AT.AccessToken)
	r.Header.Set("Authorization", "Bearer "+p.AT.AccessToken)
	r.Header.Set("Content-Type", "application/xml")
	return nil
}
