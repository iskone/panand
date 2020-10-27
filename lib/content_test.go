package lib

import (
	"fmt"
	"os"
	"testing"
)

func testnewPan() *Panand {
	return &Panand{
		Auth: nil,
		AT: &AccessToken{
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
func TestClient_AccessToken(t *testing.T) {
	c:=NewClientOauth("xfile","APP1jAFwB3P0002","0BA8A4FE89A8F33321C8AE740058D768","https://api.ihx.me/oa/caiyun")
	c.UseWap(false)
	fmt.Println(c.Login("test"))
//	fmt.Println(c.RefreshToken("35F2F9D937B4ECBBC0B5C8F2B1AC3E97A78FECB1446745D454CF090F2CF0BDA23239FD7428702853F78341E0EB4C3F1898FA809CDEED9A7488492FE8FB4A20B0742CF92FFD4C60EBF3F7051326877100"))
}
