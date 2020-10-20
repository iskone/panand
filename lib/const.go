package lib
const (
	OauthScope = "nd_cloud"
	OauthDefaultHost = "https://caiyun.feixin.10086.cn/authorize.jsp"
	OauthWapHost = "http://caiyun.feixin.10086.cn:7070/portal/oauth2"
	PanApi = "https://ose.caiyun.feixin.10086.cn"
	GetTokenApi = "/oauthApp/OAuth2/getToken"
	RefreshTokenApi = "/oauthApp/OAuth2/refreshToken"
	UserApi = "/richlifeApp/devapp/IUser"
)
const ThirdPartyAnonymousAccount = "thirdparty_anonymous_account"
const(
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
const(
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
