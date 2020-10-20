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


type ContentInfoRes struct {
	ContentInfo
}

func (p Panand) GetContentInfo() {

}
