package lib

import (
	"fmt"
	"os"
	"testing"
)

func testnewPan() *Panand {
	return &Panand{
		auth: nil,
		AccessToken: AccessToken{
			RToken: "",
			RefreshToken: RefreshToken{
				Token:    Token{ExpiresIn: 2550064505, AccessToken: os.Getenv("at")},
				OauthErr: OauthErr{},
			},
		},
	}
}
func Test_getContentInfo(t *testing.T) {
	p := testnewPan()
	r, e := p.GetContentInfo(GetContentInfoReq{
		ContentID:   "0G11ZGK0d03Y0672020101900342637k",
		OwnerMSISDN: "thirdparty_anonymous_account",
	})
	fmt.Println(r, e)
}
func TestPanand_DelCatalogContent(t *testing.T) {
	p := testnewPan()
	fmt.Println(p.DelCatalogContent(DelCatalogContent{
		ContentIDs: IDs{ID: []string{"0G11ZGK0d03Y0672020101900342637k"}},
		OprReason:  0,
	}))
}
