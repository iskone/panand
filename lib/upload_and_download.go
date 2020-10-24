package lib

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/iskone/itools"
	ixml "github.com/iskone/itools/xml"
	"io"
	"net/http"
	"strings"
)

type UploadContentList struct {
	Length            int                 `xml:"length,attr" autoLen:"UploadContentInfo"`
	UploadContentInfo []UploadContentInfo `xml:"uploadContentInfo"`
}
type WebUploadFileRequest struct {
	XMLName           xml.Name          `xml:"webUploadFileRequest"`
	OwnerMSISDN       string            `xml:"ownerMSISDN"`
	FileCount         int               `xml:"fileCount"`
	TotalSize         int               `xml:"totalSize"`
	UploadContentList UploadContentList `xml:"uploadContentList"`
	NewCatalogName    string            `xml:"newCatalogName"`
	ParentCatalogID   string            `xml:"parentCatalogID"`
}

func (p Panand) WebUploadFileRequest(upList WebUploadFileRequest) (*UploadResult, error) {
	up := itools.AutoSetLength(upList, "autoLen").(WebUploadFileRequest)
	up.OwnerMSISDN = ThirdPartyAnonymousAccount
	up.FileCount = up.UploadContentList.Length
	b, err := ixml.EncodeXml(up)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest(http.MethodPost, PanApi+IUploadAndDownloadApi, b)
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
	var readData struct {
		Result
		UploadResult UploadResult `xml:"uploadResult"`
	}
	err = ixml.DecodeXml(res.Body, &readData)
	if err != nil {
		return nil, err
	}
	return &readData.UploadResult, readData.HasErr()
}

type DownloadRequest struct {
	ContentID           string `xml:"contentID"`
	OwnerMSISDN         string `xml:"OwnerMSISDN"`
	EntryShareCatalogID string `xml:"entryShareCatalogID"`
	Operation           IBool  `xml:"operation"`
	FileVersion         int    `xml:"fileVersion"`
	Path                string `xml:"path"`
}
type DownloadUrl struct {
	ID  string
	Url string
}

func (p Panand) DownloadRequest(request DownloadRequest) ([]DownloadUrl, error) {
	var t struct {
		XMLName xml.Name `xml:"downloadRequest"`
		AppName string   `xml:"appName"`
		MSISDN  string   `xml:"MSISDN"`
		DownloadRequest
	}
	t.AppName = p.auth.Name()
	t.MSISDN = ThirdPartyAnonymousAccount
	t.DownloadRequest = request
	b, err := ixml.EncodeXml(t)
	if err != nil {
		return nil, err
	}
	fmt.Println(b.String())
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, PanApi+IUploadAndDownloadApi, b)
	if err != nil {
		return nil, err
	}
	err = p.addHeader(req)
	if err != nil {
		return nil, err
	}
	var res *http.Response
	res, err = getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	var readData struct {
		Result
		String string `xml:"String"`
	}

	err = ixml.DecodeXml(res.Body, &readData)

	if err != nil {
		return nil, err
	}
	if err = readData.HasErr(); err != nil {
		return nil, err
	}
	var r []DownloadUrl
	fmt.Println(readData.String)
	if !t.Operation.bool {
		r = append(r, DownloadUrl{
			ID:  t.ContentID,
			Url: readData.String,
		})
	} else {
		var v struct {
			XMLName xml.Name `xml:"item"`
			Text    string   `xml:",chardata"`
		}

		dec := xml.NewDecoder(bytes.NewBufferString(readData.String))
		dec.Strict = false
		for {
			err = dec.Decode(&v)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			s := strings.Split(v.Text, "|")
			if len(s) != 2 {
				continue
			}
			r = append(r, DownloadUrl{
				ID:  s[0],
				Url: s[1],
			})
		}

	}
	return r, nil
}

func (p Panand) PcUploadFileRequest(parentCatalogID string, uploadContentList []UploadContentInfo, totalSize int) (*UploadResult, error) {
	var reqStruct struct {
		XMLName           xml.Name          `xml:"pcUploadFileRequest"`
		OwnerMSISDN       string            `xml:"ownerMSISDN"`
		FileCount         int               `xml:"fileCount"`
		TotalSize         int               `xml:"totalSize"`
		UploadContentList UploadContentList `xml:"uploadContentList"`
		NewCatalogName    string            `xml:"newCatalogName"`
		ParentCatalogID   string            `xml:"parentCatalogID"`
	}
	reqStruct.OwnerMSISDN = ThirdPartyAnonymousAccount
	reqStruct.FileCount = len(uploadContentList)
	reqStruct.TotalSize = totalSize
	reqStruct.UploadContentList.UploadContentInfo = uploadContentList
	reqStruct.UploadContentList.Length = reqStruct.FileCount
	reqStruct.ParentCatalogID = parentCatalogID
	b, err := ixml.EncodeXml(reqStruct)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, PanApi+IUploadAndDownloadApi, b)
	if err != nil {
		return nil, err
	}
	err = p.addHeader(req)
	if err != nil {
		return nil, err
	}
	var res *http.Response
	res, err = getHttpClient().Do(req)
	if err != nil {
		return nil, err
	}
	var readData struct {
		Result
		UploadResult UploadResult `xml:"uploadResult"`
	}
	err = ixml.DecodeXml(res.Body, &readData)
	if err != nil {
		return nil, err
	}
	return &readData.UploadResult, readData.HasErr()
}
