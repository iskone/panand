package lib

import "encoding/xml"

type SimpleSearchReq struct {
	Text     string               `xml:",chardata"`
	Account  string               `xml:"account"`
	Scope    SimpleSearchReqScope `xml:"scope"`
	Keyword  string               `xml:"keyword"`
	StartNum int                  `xml:"startNum"`
	EndNum   int                  `xml:"endNum"`
}
type SimpleSearch struct {
	XMLName         xml.Name        `xml:"simpleSearch"`
	Text            string          `xml:",chardata"`
	SimpleSearchReq SimpleSearchReq `xml:"simpleSearchReq"`
}
type SimpleSearchReqScope string

const SimpleSearchReqScopeFile SimpleSearchReqScope = "X1"
const SimpleSearchReqScopeDir SimpleSearchReqScope = "X2"

func (p Panand) SimpleSearch(scope SimpleSearchReqScope, keyword string, start, end int) {
	var ss SimpleSearch
	ss.SimpleSearchReq.Account = ThirdPartyAnonymousAccount
	ss.SimpleSearchReq.Scope = scope
	ss.SimpleSearchReq.Keyword = keyword
	ss.SimpleSearchReq.StartNum = start
	ss.SimpleSearchReq.EndNum = end

}
