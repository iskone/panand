package lib

import (
	"encoding/xml"
	"time"
)

type IDs struct {
	ID []string `xml:"ID"`
}
type Err struct {
	Code string
	Desc string
}
func (e Err) Error() string {
	msg := "code: " + e.Code
	if e.Desc != "" {
		msg += ",desc: " + e.Desc

	}
	return msg
}
type Result struct {
	XMLName    xml.Name `xml:"result"`
	Text       string   `xml:",chardata"`
	ResultCode string   `xml:"resultCode,attr"`
	Desc       string   `xml:"desc,attr"`
	RetCode    string   `xml:"retCode"`
}
func (r Result) HasErr() error {
	e := Err{
		Desc: r.Desc,
	}
	if r.RetCode != "0" && r.RetCode != "" {
		e.Code = r.RetCode
		return e
	}
	if r.ResultCode != "0" && r.ResultCode != "" {
		e.Code = r.ResultCode
		return e
	}
	return nil
}

type ITime struct {
	time.Time
}

func (I ITime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//YYYYMMDDHH24MMSS
	s:=I.Format("20060102150405")
	return  e.EncodeElement(s,start)

}

func (I *ITime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*I = ITime{}
	var s string
	err:=d.Decode(&s)
	if err != nil {
		return err
	}
	I.Time,err = time.Parse("20060102150405",s)
	return err
}

type contentType int
type contentOrigin int
type safeState int
type transferState int
type ETagOprType int
type auditResult int
type ContentInfo struct {
	XMLName       xml.Name      `xml:"contentInfo"`
	Text          string        `xml:",chardata"`
	ContentID     string        `xml:"contentID"`
	ContentName   string        `xml:"contentName"`
	ContentSuffix string        `xml:"contentSuffix"`
	ContentSize   int           `xml:"contentSize"`
	ContentDesc   string        `xml:"contentDesc"`
	IsShared      bool          `xml:"isShared"`
	ContentType   contentType   `xml:"contentType"`
	ContentOrigin contentOrigin `xml:"contentOrigin"`

	UpdateTime   string `xml:"updateTime"`
	CommentCount int    `xml:"commentCount"`

	ThumbnailURL    string        `xml:"thumbnailURL"`
	BigThumbnailURL string        `xml:"bigthumbnailURL"`
	PresentURL      string        `xml:"presentURL"`
	PresentLURL     string        `xml:"presentLURL"`
	PresentHURL     string        `xml:"presentHURL"`
	ShareDoneeCount int           `xml:"shareDoneeCount"`
	SafeState       safeState     `xml:"safestate"`
	TransferState   transferState `xml:"transferstate"`
	IsFocusContent  bool          `xml:"isFocusContent"`

	UpdateShareTime string `xml:"updateShareTime"`
	UploadTime      string `xml:"uploadTime"`

	ETagOprType ETagOprType `xml:"ETagOprType"`

	OpenType        bool        `xml:"openType"`
	AuditResult     auditResult `xml:"auditResult"`
	ParentCatalogId string      `xml:"parentCatalogId"`
	Channel         string      `xml:"channel"`
	GeoLocFlag      string      `xml:"geoLocFlag"`
	Digest          string      `xml:"digest"`
	Version         int         `xml:"version"`
	FileEtag        int         `xml:"fileEtag"`
	FileVersion     int         `xml:"fileVersion"`
	TombStoned      bool        `xml:"tombstoned"`
	ProxyID         string      `xml:"proxyID"`
	Moved           bool        `xml:"moved"`
	MidThumbnailURL string      `xml:"midthumbnailURL"`
	Owner           string      `xml:"owner"`
	Modifier        string      `xml:"modifier"`
	ShareType       int         `xml:"shareType"`
}
type LiteContentInfo struct {
	XMLName       xml.Name      `xml:"liteContentInfo"`
	Text          string        `xml:",chardata"`
	ContentID     string        `xml:"contentID"`
	FileEtag        int         `xml:"fileEtag"`
}
type CatalogInfo struct {
	XMLName       xml.Name      `xml:"catalogInfo"`
	Text          string        `xml:",chardata"`
	CatalogID          string        `xml:"catalogID"`
	atalogName          string        `xml:"catalogName"`
}
