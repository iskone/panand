package lib

import (
	"encoding/xml"
	"errors"
	"github.com/iskone/itools"
	iXml "github.com/iskone/itools/xml"
	"net/http"
)

type GetContentInfoReq struct {
	XMLName             xml.Name `xml:"getContentInfo"`
	ContentID           string   `xml:"contentID"`
	EntryShareCatalogID string   `xml:"entryShareCatalogID"`
	OwnerMSISDN         string   `xml:"ownerMSISDN"`
	Path                string   `xml:"path,omitempty"`
}

type DelCatalogContent struct {
	XMLName     xml.Name `xml:"delCatalogContent"`
	OwnerMSISDN string   `xml:"ownerMSISDN"`
	CatalogIDs  IDs      `xml:"catalogIDs"`
	ContentIDs  IDs      `xml:"contentIDs"`
	//ContentURLs []string `xml:"contentURLs"`
	OprReason int `xml:"oprReason"`
}

func (p Panand) GetContentInfo(gReq GetContentInfoReq) (*ContentInfo, error) {
	mReq := gReq
	mReq.OwnerMSISDN = ThirdPartyAnonymousAccount
	bs, e := iXml.EncodeXml(mReq)
	if e != nil {
		return nil, e
	}
	req, err := http.NewRequest(http.MethodPost, PanApi+IContentApi, bs)
	if err != nil {
		return nil, err
	}
	if err = p.addHeader(req); err != nil {
		return nil, err
	}
	res, err := getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	type ContentInfoRes struct {
		Result
		ContentInfo ContentInfo `xml:"contentInfo"`
	}
	var cir ContentInfoRes
	err = iXml.DecodeXml(res.Body, &cir)
	if err != nil {
		return nil, err
	}
	if err = cir.HasErr(); err != nil {
		return nil, err
	}
	return &cir.ContentInfo, nil
}
func (p Panand) DelCatalogContent(del DelCatalogContent) error {
	md := itools.AutoSetLength(del, "autoLen").(DelCatalogContent)
	if md.CatalogIDs.Length > 200 {
		return errors.New("CatalogIDs > 200")
	}
	if md.ContentIDs.Length > 200 {
		return errors.New("ContentID > 200")
	}
	if md.ContentIDs.Length+md.CatalogIDs.Length > 200 {
		return errors.New("ContentID + CatalogIDs > 200")
	}
	md.OwnerMSISDN = ThirdPartyAnonymousAccount
	if md.OprReason != 0 {
		md.OprReason = 1
	}
	bs, e := iXml.EncodeXml(md)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(http.MethodPost, PanApi+IContentApi, bs)
	if e != nil {
		return e
	}
	if e = p.addHeader(req); e != nil {
		return e
	}
	res, e := getHttpClient().Do(req)
	if e != nil {
		return e
	}
	var r Result
	if e = iXml.DecodeXml(res.Body, &r); e != nil {
		return e
	}
	return r.HasErr()
}
