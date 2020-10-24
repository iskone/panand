package lib

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/iskone/itools"
	ixml "github.com/iskone/itools/xml"
	"net/http"
)

type GetDiskReq struct {
	CatalogID           string          `xml:"catalogID"`
	EntryShareCatalogID string          `xml:"entryShareCatalogID"`
	FilterType          filterType      `xml:"filterType"`
	CatalogSortType     catalogSortType `xml:"catalogSortType"`
	ContentType         contentType     `xml:"contentType"`
	ContentSortType     contentSortType `xml:"contentSortType"`
	SortDirection       IBool           `xml:"sortDirection"`
	StartNumber         int             `xml:"startNumber"`
	EndNumber           int             `xml:"endNumber"`
	ChannelList         string          `xml:"channelList"`
	CatalogType         catalogType     `xml:"catalogType"`
	Path                string          `xml:"path"`
}

func (p Panand) GetDisk(gdr GetDiskReq) (*GetDiskResult, error) {
	type hr struct {
		XMLName xml.Name `xml:"getDisk"`
		MSISDN  string   `xml:"MSISDN"`
		GetDiskReq
	}
	rhr := hr{
		MSISDN:     ThirdPartyAnonymousAccount,
		GetDiskReq: gdr,
	}
	b, err := ixml.EncodeXml(rhr)
	fmt.Println(b.String())
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, PanApi+ICatalogApi, b)
	if err = p.addHeader(req); err != nil {
		return nil, err
	}
	res, err := getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	type resT struct {
		Result
		GetDiskResult GetDiskResult `xml:"getDiskResult"`
	}
	var gd resT
	if err = ixml.DecodeXml(res.Body, &gd); err != nil {
		return nil, err
	}
	if err = gd.HasErr(); err != nil {
		return nil, err
	}
	fmt.Println(gd.GetDiskResult)
	r := itools.AutoSetLength(gd.GetDiskResult, "autoLen").(GetDiskResult)
	return &r, err
}

type CreateCatalogExtReq struct {
	ParentCatalogID string      `xml:"parentCatalogID"`
	NewCatalogName  string      `xml:"newCatalogName"`
	CatalogType     catalogType `xml:"catalogType"`
	Path            string      `xml:"path"`
	ManualRename    int         `xml:"manualRename"`
}

func (p Panand) CreatCatalogExt(ccx CreateCatalogExtReq) (*CatalogInfo, error) {
	type ReqStruct struct {
		XMLName             xml.Name `xml:"createCatalogExt"`
		CreateCatalogExtReq struct {
			MSISDN string `xml:"ownerMSISDN"`
			CreateCatalogExtReq
		} `xml:"createCatalogExtReq"`
	}
	var xreq ReqStruct
	xreq.CreateCatalogExtReq.MSISDN = ThirdPartyAnonymousAccount
	xreq.CreateCatalogExtReq.CreateCatalogExtReq = ccx
	bs, err := ixml.EncodeXml(xreq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, PanApi+ICatalogApi, bs)
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
	var getres struct {
		Result
		CatalogInfo CatalogInfo `xml:"catalogInfo"`
	}
	err = ixml.DecodeXml(res.Body, &getres)
	if err != nil {
		return nil, err
	}
	return &getres.CatalogInfo, getres.HasErr()
}

func (p Panand) CopyContentCatalog(newCatalogID string, contentInfoList, catalogInfoList []string) ([]string, error) {
	if len(catalogInfoList) > 50 {
		return nil, errors.New("catalogInfoList > 50")
	}
	if len(contentInfoList) > 50 {
		return nil, errors.New("contentInfoList > 50")
	}
	var reqData struct {
		XMLName         xml.Name `xml:"copyContentCatalog"`
		MSISDN          string   `xml:"MSISDN"`
		ContentInfoList IDs      `xml:"contentInfoList"`
		CatalogInfoList IDs      `xml:"catalogInfoList"`
		NewCatalogID    string   `xml:"newCatalogID"`
	}
	reqData.MSISDN = ThirdPartyAnonymousAccount
	reqData.NewCatalogID = newCatalogID
	reqData.CatalogInfoList = IDs{
		Length: len(catalogInfoList),
		ID:     catalogInfoList,
	}
	reqData.ContentInfoList = IDs{
		Length: len(contentInfoList),
		ID:     contentInfoList,
	}
	b, err := ixml.EncodeXml(reqData)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest(http.MethodPost, PanApi+ICatalogApi, b)
	if err != nil {
		return nil, err
	}
	err = p.addHeader(req)
	if err != nil {
		return nil, err
	}
	res, err = getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	var read struct {
		Result
		Array struct {
			Text   string   `xml:",chardata"`
			Length string   `xml:"length,attr"`
			Item   []string `xml:"item"`
		} `xml:"array"`
	}
	err = ixml.DecodeXml(res.Body, &read)
	if err != nil {
		return nil, err
	}
	return read.Array.Item, read.HasErr()

}

func (p Panand) UpdateCatalogInfo(catalogID, catalogName string, ctype catalogType, path string) (string, error) {
	var reqData struct {
		XMLName     xml.Name    `xml:"updateCatalogInfo"`
		MSISDN      string      `xml:"MSISDN"`
		CatalogID   string      `xml:"catalogID"`
		CatalogName string      `xml:"catalogName"`
		CatalogType catalogType `xml:"catalogType"`
		Path        string      `xml:"path"`
	}
	reqData.MSISDN = ThirdPartyAnonymousAccount
	reqData.CatalogID = catalogID
	reqData.CatalogName = catalogName
	reqData.CatalogType = ctype
	reqData.Path = path
	b, err := ixml.EncodeXml(reqData)
	if err != nil {
		return "", err
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, PanApi+ICatalogApi, b)
	if err != nil {
		return "", err
	}
	err = p.addHeader(req)
	if err != nil {
		return "", err
	}
	var res *http.Response
	res, err = getHttpClient().Do(req)
	if err != nil {
		return "", err
	}
	var readRes struct {
		Result
		UpdateCatalogInfoRes struct {
			Text    string `xml:",chardata"`
			DirEtag string `xml:"dirEtag"`
		} `xml:"updateCatalogInfoRes"`
	}
	err = ixml.DecodeXml(res.Body, &readRes)
	if err != nil {
		return "", err
	}
	return readRes.UpdateCatalogInfoRes.DirEtag, readRes.HasErr()
}

type MoveContentCatalogRes struct {
	LiteContentInfoList struct {
		Text            string `xml:",chardata"`
		Length          string `xml:"length,attr"`
		LiteContentInfo LiteContentInfo `xml:"liteContentInfo"`
	} `xml:"liteContentInfoList"`
	LiteCatalogInfoList struct {
		Text            string `xml:",chardata"`
		Length          string `xml:"length,attr"`
		LiteCatalogInfo LiteCatalogInfo `xml:"liteCatalogInfo"`
	} `xml:"liteCatalogInfoList"`
}
func (p Panand) MoveContentCatalog (newCatalogID string, contentInfoList, catalogInfoList []string) (*MoveContentCatalogRes, error) {
	if len(catalogInfoList) > 200 {
		return nil, errors.New("catalogInfoList > 200")
	}
	if len(contentInfoList) > 200 {
		return nil, errors.New("contentInfoList > 200")
	}
	var reqData struct {
		XMLName         xml.Name `xml:"moveContentCatalog"`
		MSISDN          string   `xml:"MSISDN"`
		ContentInfoList IDs      `xml:"contentInfoList"`
		CatalogInfoList IDs      `xml:"catalogInfoList"`
		NewCatalogID    string   `xml:"newCatalogID"`
	}
	reqData.MSISDN = ThirdPartyAnonymousAccount
	reqData.NewCatalogID = newCatalogID
	reqData.CatalogInfoList = IDs{
		Length: len(catalogInfoList),
		ID:     catalogInfoList,
	}
	reqData.ContentInfoList = IDs{
		Length: len(contentInfoList),
		ID:     contentInfoList,
	}
	b, err := ixml.EncodeXml(reqData)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest(http.MethodPost, PanApi+ICatalogApi, b)
	if err != nil {
		return nil, err
	}
	err = p.addHeader(req)
	if err != nil {
		return nil, err
	}
	res, err = getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	var read struct {
		Result
		MoveContentCatalogRes MoveContentCatalogRes `xml:"moveContentCatalogRes"`
	}
	err = ixml.DecodeXml(res.Body, &read)
	if err != nil {
		return nil, err
	}
	return &read.MoveContentCatalogRes, read.HasErr()

}
