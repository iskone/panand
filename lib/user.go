package lib

import (
	"encoding/xml"
	ixml "github.com/iskone/itools/xml"
"net/http"
)

type GetDiskInfo struct {
	XMLName xml.Name `xml:"getDiskInfo"`
	Text    string   `xml:",chardata"`
	MSISDN  string   `xml:"MSISDN"`
}
type DiskInfo struct {
	FreeDiskSize int `xml:"freeDiskSize"`
	DiskSize     int `xml:"diskSize"`
}

func (p Panand) GetDiskInfo() (DiskInfo, error) {

	buf, _ := ixml.EncodeXml(GetDiskInfo{
		MSISDN: "thirdparty_anonymous_account",
	})
	type ResStruct struct {
		Result
		DiskInfo DiskInfo `xml:"diskInfo"`
	}
	req, e := http.NewRequest(http.MethodPost, PanApi+UserApi, buf)
	if e != nil {
		return DiskInfo{}, e
	}
	p.addHeader(req)
	res, e := getHttpClient().Do(req)
	if e != nil {
		return DiskInfo{}, e
	}
	var resData ResStruct
	if e = ixml.DecodeXml(res.Body, &resData); e != nil {
		return DiskInfo{}, e
	}
	if e = resData.HasErr(); e != nil {
		return DiskInfo{}, e
	}
	return resData.DiskInfo, nil
}
