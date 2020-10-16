package lib

import "encoding/xml"

type ContentInfoReq struct {
	XMLName             xml.Name `xml:"getContentInfo"`
	ContentID           string   `xml:"contentID"`
	EntryShareCatalogID string   `xml:"entryShareCatalogID"`
	OwnerMSISDN         string   `xml:"ownerMSISDN"`
	Path                string   `xml:"path,omitempty"`
}
type DelCatalogContent struct {
	XMLName     xml.Name `xml:"delCatalogContent"`
	OwnerMSISDN string   `xml:"ownerMSISDN"`
	CatalogIDs  struct {
		ID     []string `xml:"ID"`
	} `xml:"catalogIDs"`
	ContentIDs IDs `xml:"contentIDs"`
	ContentURLs []string `xml:"contentURLs"`
	OprReason string `xml:"oprReason"`
}

type ContentInfo struct {
	XMLName         xml.Name       `xml:"contentInfo"`
	Text            string         `xml:",chardata"`
	ContentID       string         `xml:"contentID"`
	ContentName     string         `xml:"contentName"`
	ParentCatalogId string         `xml:"parentCatalogId"`
	ContentSize     string         `xml:"contentSize"`
	ContentDesc     string         `xml:"contentDesc"`
	ContentType     string         `xml:"contentType"`
	IsShared        string         `xml:"isShared"`
	ThumbnailURL    string         `xml:"thumbnailURL"`
	UpdateTime      string         `xml:"updateTime"`
	ContentOrigin   string         `xml:"contentOrigin"`
	Safestate       string         `xml:"safestate"`
	BigthumbnailURL string         `xml:"bigthumbnailURL"`
	PresentURL      string         `xml:"presentURL"`
	CommentCount    string         `xml:"commentCount"`
	ContentTAGList  ContentTAGList `xml:"contentTAGList"`
	UploadTime      string         `xml:"uploadTime"`
	ShareDoneeCount string         `xml:"shareDoneeCount"`
	IsFocusContent  string         `xml:"isFocusContent"`
	ETagOprType     string         `xml:"ETagOprType"`
	ContentSuffix   string         `xml:"contentSuffix"`
	Transferstate   string         `xml:"transferstate"`
	OpenType        string         `xml:"openType"`
	AuditResult     string         `xml:"auditResult"`
	Channel         string         `xml:"channel"`
	GeoLocFlag      string         `xml:"geoLocFlag"`
	Digest          string         `xml:"digest"`
	FileEtag        string         `xml:"fileEtag"`
}
type ContentTAGList struct {
	XMLName xml.Name `xml:"contentTAGList"`
	Text    string   `xml:",chardata"`
	Length  string   `xml:"length,attr"`
}
type ContentInfoRes struct {
	ContentInfo
}

func (p Panand) GetContentInfo() {

}
