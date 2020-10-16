package lib

import (
	"fmt"
	"net/http"
	"time"
)

type Panand struct {
	auth Auth
	AccessToken
}

func (p Panand) addHeader(r *http.Request) error {
	if p.ExpiresIn <= int(time.Now().Unix()) {
		if rt,e:=p.auth.RefreshToken(p.RefreshToken);e!=nil{
			return e
		} else {
			p.Token = rt.Token
		}
	}
	fmt.Println(p.AccessToken.AccessToken)
	r.Header.Set("Authorization","Bearer "+p.AccessToken.AccessToken)
	r.Header.Set("Content-Type", "application/xml")
	return nil
}
