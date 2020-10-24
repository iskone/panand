package lib

const (
	OauthScope       = "nd_cloud"
	OauthDefaultHost = "https://caiyun.feixin.10086.cn/authorize.jsp"
	OauthWapHost     = "http://caiyun.feixin.10086.cn:7070/portal/oauth2"
	PanApi           = "https://ose.caiyun.feixin.10086.cn"
	GetTokenApi      = "/oauthApp/OAuth2/getToken"
	RefreshTokenApi  = "/oauthApp/OAuth2/refreshToken"
	UserApi          = "/richlifeApp/devapp/IUser"
	IContentApi      = "/richlifeApp/devapp/IContent"
	ICatalogApi      = "/richlifeApp/devapp/ICatalog"
	IUploadAndDownloadApi      = "/richlifeApp/devapp/IUploadAndDownload"
)
const ThirdPartyAnonymousAccount = "thirdparty_anonymous_account"
const (
	ContentTypeOther contentType = iota
	ContentTypeImg
	ContentTypeAudio
	ContentTypeVideo
	ContentTypeMessage
	ContentTypeWord
	ContentTypeShare
	ContentTypeExel
	ContentTypePpt
)
const (
	ContentOriginPCPush contentOrigin = iota
	ContentOriginPCClient
	ContentOriginMobile
	ContentOriginShare
	ContentOrigin17Play
	ContentOriginSuperEmail
	ContentOriginHome
	ContentOriginMessage
	ContentOriginPhoneEmail
	ContentOriginWeb
	ContentOriginCommunity
	ContentOriginOther contentOrigin = 99
)
const (
	SafeStateDefault safeState = iota
	SafeStateSafe
	SafeStateVirus
	SafeStateCheckIng
	SafeStateUnknown
)
const (
	TransferStateIng transferState = iota
	TransferStateFail
	TransferStateExamine
	TransferStateNormal
	TransferStateShield
)
const (
	ETagOprTypeModify ETagOprType = iota
	ETagOprTypeAdded
	ETagOprTypeRemove
)
const (
	AuditDefault auditResult = iota
	AuditPass
	AuditFail
	AuditManual
	AuditExamine
)

const (
	CatalogTypeNotUpdate catalogType = -2
	CatalogTypeAll       catalogType = -1
	CatalogTypeCommon    catalogType = iota
	CatalogTypeImage
	CatalogTypeAudio
	CatalogTypeVideo
	CatalogTypeMessage
	CatalogTypeDocument
	CatalogTypeApp
	CatalogTypeSync
	CatalogTypeVirtual
)
const (
	OpenTypePrivate openType = iota
	OpenTypePublic
)
const (
	ShareTypeNotShare shareType = iota
	ShareTypeExternal
	ShareTypeP2P
	ShareTypeExternalAndP2P
	ShareTypeEnt
)
const (
	ShareFlagDefault shareFlag = iota
	ShareFlagSendShare
	ShareFlagReceiveShare
	ShareFlagSendCooperation
	ShareFlagReceiveCooperation
)
const (
	FilterTypeAll filterType = iota
	FilterTypeCatalog
	FilterTypeContent
)
const (
	CatalogSortTime catalogSortType = iota
	CatalogSortName
)
const (
	ContentSortTime contentSortType = iota
	ContentSortName
	ContentSortSize
)
