package lib

import (
	"encoding/xml"
	"time"
)

type IDs struct {
	Length int      `xml:"length,attr" autoLen:"ID"`
	ID     []string `xml:"ID"`
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
	s := I.Format("20060102150405")
	return e.EncodeElement(s, start)
}

func (I *ITime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*I = ITime{}
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	I.Time, err = time.Parse("20060102150405", s)
	return err
}

type IBool struct {
	bool
}

func (I IBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//YYYYMMDDHH24MMSS
	i := 0
	if I.bool {
		i = 1
	}
	return e.EncodeElement(i, start)
}

func (I *IBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*I = IBool{}
	var b bool
	err := d.DecodeElement(&b, &start)
	if err != nil {
		return err
	}
	I.bool = b
	return err
}

type filterType int
type contentType int
type contentOrigin int
type safeState int
type transferState int
type ETagOprType int
type auditResult int
type catalogType int
type openType int
type shareType int
type shareFlag int
type catalogSortType int
type contentSortType int
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

	UpdateTime   ITime `xml:"updateTime"`
	CommentCount int   `xml:"commentCount"`

	ThumbnailURL    string        `xml:"thumbnailURL"`
	BigThumbnailURL string        `xml:"bigthumbnailURL"`
	PresentURL      string        `xml:"presentURL"`
	PresentLURL     string        `xml:"presentLURL"`
	PresentHURL     string        `xml:"presentHURL"`
	ShareDoneeCount int           `xml:"shareDoneeCount"`
	SafeState       safeState     `xml:"safestate"`
	TransferState   transferState `xml:"transferstate"`
	IsFocusContent  bool          `xml:"isFocusContent"`

	UpdateShareTime ITime `xml:"updateShareTime"`
	UploadTime      ITime `xml:"uploadTime"`

	ETagOprType ETagOprType `xml:"ETagOprType"`

	OpenType        openType    `xml:"openType"`
	AuditResult     auditResult `xml:"auditResult"`
	ParentCatalogId string      `xml:"parentCatalogId"`
	Channel         string      `xml:"channel"`
	GeoLocFlag      string      `xml:"geoLocFlag"`
	Digest          string      `xml:"digest"`
	Version         int         `xml:"version"`
	FileEtag        int         `xml:"fileEtag"`
	FileVersion     int         `xml:"fileVersion"`
	TombStoned      IBool       `xml:"tombstoned"`
	ProxyID         string      `xml:"proxyID"`
	Moved           IBool       `xml:"moved"`
	MidThumbnailURL string      `xml:"midthumbnailURL"`
	Owner           string      `xml:"owner"`
	Modifier        string      `xml:"modifier"`
	ShareType       shareType   `xml:"shareType"`
}
type LiteContentInfo struct {
	XMLName   xml.Name `xml:"liteContentInfo"`
	Text      string   `xml:",chardata"`
	ContentID string   `xml:"contentID"`
	FileEtag  int      `xml:"fileEtag"`
}
type CatalogInfo struct {
	XMLName         xml.Name    `xml:"catalogInfo"`
	Text            string      `xml:",chardata"`
	CatalogID       string      `xml:"catalogID"`
	CatalogName     string      `xml:"catalogName"`
	CatalogType     catalogType `xml:"catalogType"`
	CreateTime      ITime       `xml:"createTime"`
	UpdateTime      ITime       `xml:"updateTime"`
	IsShared        bool        `xml:"isShared"`
	CatalogLevel    int         `xml:"catalogLevel"`
	ShareDoneeCount int         `xml:"shareDoneeCount"`
	ETagOprType     ETagOprType `xml:"ETagOprType"`
	OpenType        openType    `xml:"openType"`
	ParentCatalogId string      `xml:"parentCatalogId"`
	DirEtag         int         `xml:"dirEtag"`
	TombStoned      IBool       `xml:"tombstoned"`
	ProxyID         string      `xml:"proxyID"`
	Moved           IBool       `xml:"moved"`
	IsFixedDir      IBool       `xml:"isFixedDir"`
	IsSynced        IBool       `xml:"isSynced"`
	Owner           string      `xml:"owner"`
	Modifier        string      `xml:"modifier"`
	ShareType       shareType   `xml:"shareType"`
}

type CatalogNode struct {
	XMLName       xml.Name `xml:"catalogNode"`
	Text          string   `xml:",chardata"`
	CatalogID     string   `xml:"catalogID"`
	CatalogName   string   `xml:"catalogName"`
	DirEtag       int      `xml:"dirEtag"`
	ChildNodeList struct {
		Length      int           `xml:"length,attr" autoLen:"CatalogNode"`
		CatalogNode []CatalogNode `xml:"catalogNode"`
	} `xml:"childNodeList"`
}

type GetDiskResult struct {
	XMLName         xml.Name `xml:"getDiskResult"`
	ParentCatalogId string   `xml:"parentCatalogId"`
	NodeCount       int   `xml:"nodeCount"`
	CatalogList     struct {
		Length      int           `xml:"length,attr" autoLen:"CatalogInfo"`
		CatalogInfo []CatalogInfo `xml:"catalogInfo"`
	} `xml:"catalogList"`
	ContentList struct {
		Length      int           `xml:"length,attr" autoLen:"ContentInfo"`
		ContentInfo []ContentInfo `xml:"contentInfo"`
	} `xml:"contentList"`
}

type UserCtnCtlg struct {
	XMLName   xml.Name `xml:"userCtnCtlg"`
	Text      string   `xml:",chardata"`
	CtnCount  int      `xml:"ctnCount"`
	CLtgCount int      `xml:"cltgCount"`
}
type LiteCatalogInfo struct {
	XMLName   xml.Name `xml:"liteCatalogInfo"`
	Text      string   `xml:",chardata"`
	CatalogID string   `xml:"catalogID"`
	DirEtag   int      `xml:"dirEtag"`
}

//=======
type UploadContentInfo struct {
	XMLName         xml.Name `xml:"uploadContentInfo"`
	Text            string   `xml:",chardata"`
	ContentName     string   `xml:"contentName"`
	ContentSize     int      `xml:"contentSize"`
	ContentDesc     string   `xml:"contentDesc"`
	ContentTAGList  string   `xml:"contentTAGList"`
	ComlexFlag      IBool    `xml:"comlexFlag"`
	ComlexCID       string   `xml:"comlexCID"`
	ResCID          IDs      `xml:"resCID"`
	Digest          string   `xml:"digest"`
	UpdateContentID string   `xml:"updateContentID"`
	FileEtag        int      `xml:"fileEtag"`
	FileVersion     int      `xml:"fileVersion"`
}

type NewContent struct {
	XMLName      xml.Name `xml:"newContent"`
	Text         string   `xml:",chardata"`
	ContentID    string   `xml:"contentID"`
	ContentName  string   `xml:"contentName"`
	IsNeedUpload IBool    `xml:"isNeedUpload"`
	FileEtag     int      `xml:"fileEtag"`
	FileVersion  int      `xml:"fileVersion"`
}

type UploadResult struct {
	XMLName          xml.Name `xml:"uploadResult"`
	Text             string   `xml:",chardata"`
	UploadTaskID     string   `xml:"uploadTaskID"`
	RedirectionUrl   string   `xml:"redirectionUrl"`
	NewContentIDList struct {
		Length     int          `xml:"length,attr" autoLen:"NewContent"`
		NewContent []NewContent `xml:"newContent"`
	} `xml:"newContentIDList"`
	CatalogIDList IDs `xml:"catalogIDList"`
}

type FileUploadInfo struct {
	XMLName   xml.Name `xml:"fileUploadInfo"`
	Text      string   `xml:",chardata"`
	ContentID string   `xml:"contentID"`
	FName     string   `xml:"fName"`
	Pgs       string   `xml:"pgs"`
}
type UploadTaskInfo struct {
	XMLName         xml.Name `xml:"uploadTaskInfo"`
	Text            string   `xml:",chardata"`
	TaskID          string   `xml:"taskID"`
	UploadURL       string   `xml:"uploadURL"`
	FileUploadInfos struct {
		Length         int              `xml:"length,attr" autoLen:"FileUploadInfo"`
		FileUploadInfo []FileUploadInfo `xml:"fileUploadInfo"`
	} `xml:"fileUploadInfos"`
}
type LiteTaskInfo struct {
	XMLName   xml.Name `xml:"liteTaskInfo"`
	Text      string   `xml:",chardata"`
	TaskID    string   `xml:"taskID"`
	ContentID string   `xml:"contentID"`
	Path      string   `xml:"path"`
}

//==========
type DiskInfo struct {
	XMLName      xml.Name `xml:"diskInfo"`
	FreeDiskSize int      `xml:"freeDiskSize"`
	DiskSize     int      `xml:"diskSize"`
}

//======
type ShareInfo struct {
	XMLName   xml.Name  `xml:"shareInfo"`
	Text      string    `xml:",chardata"`
	ShareFlag shareFlag `xml:"shareFlag"`
	ReadFlag  IBool     `xml:"readFlag"`
	ShareTime ITime     `xml:"shareTime"`
	Owner     string    `xml:"owner"`
	Donee     string    `xml:"donee"`
	CmntFlag  IBool     `xml:"cmntFlag"`
	DlFlag    IBool     `xml:"dlFlag"`
}
type SrchCtlgInfo struct {
	XMLName     xml.Name    `xml:"srchCtlgInfo"`
	Text        string      `xml:",chardata"`
	CatalogInfo CatalogInfo `xml:"catalogInfo"`
	ShareInfo   ShareInfo   `xml:"shareInfo"`
}
type SrchCntInfo struct {
	XMLName     xml.Name    `xml:"srchCntInfo"`
	Text        string      `xml:",chardata"`
	ContentInfo CatalogInfo `xml:"contentInfo"`
	ShareInfo   ShareInfo   `xml:"shareInfo"`
}

//===

type GetOutLinkResOne struct {
	XMLName xml.Name `xml:"getOutLinkResOne"`
	Text    string   `xml:",chardata"`
	ObjID   string   `xml:"objID"`
	LinkID  string   `xml:"linkID"`
	LinkUrl string   `xml:"linkUrl"`
}
type OutLink struct {
	XMLName     xml.Name    `xml:"outLink"`
	Text        string      `xml:",chardata"`
	LinkID      string      `xml:"linkID"`
	Url         string      `xml:"url"`
	LkName      string      `xml:"lkName"`
	CtTime      ITime       `xml:"ctTime"`
	LastUdTime  ITime       `xml:"lastUdTime"`
	LastAudTime ITime       `xml:"lastAudTime"`
	FNum        int         `xml:"fNum"`
	TolSize     int         `xml:"tolSize"`
	AudRes      auditResult `xml:"audRes"`
	AudDesc     string      `xml:"audDesc"`
	PubType     int         `xml:"pubType"`
	DlTimes     int         `xml:"dlTimes"`
}

type OutLinkCaInfo struct {
	XMLName xml.Name `xml:"outLinkCaInfo"`
	Text    string   `xml:",chardata"`
	CaID    string   `xml:"caID"`
	CaName  string   `xml:"caName"`
	CtTime  ITime    `xml:"ctTime"`
	UdTime  ITime    `xml:"udTime"`
	DlTimes int      `xml:"dlTimes"`
}
type OutLinkCoInfo struct {
	XMLName       xml.Name      `xml:"outLinkCoInfo"`
	Text          string        `xml:",chardata"`
	CaID          string        `xml:"caID"`
	CaName        string        `xml:"caName"`
	CoSuffix      string        `xml:"coSuffix"`
	CoSize        int           `xml:"coSize"`
	UdTime        ITime         `xml:"udTime"`
	ThumbnailURL  string        `xml:"thumbnailURL"`
	BThumbnailURL string        `xml:"bthumbnailURL"`
	PresentURL    string        `xml:"presentURL"`
	SafeState     safeState     `xml:"safestate"`
	AuditResult   auditResult   `xml:"auditResult"`
	TransferState transferState `xml:"transferstate"`
	DlTimes       int           `xml:"dlTimes"`
}
