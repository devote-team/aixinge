package validation

var (
	Id                     = Rules{"ID": {NotEmpty()}}
	Api                    = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	Menu                   = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMeta               = Rules{"Title": {NotEmpty()}}
	Login                  = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	UserCreate             = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfo               = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	PurchasePageInfo       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "PurchaseOrderId": {NotEmpty()}}
	SalesPageInfo          = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "SalesOrderId": {NotEmpty()}}
	PurchaseReturnPageInfo = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "PurchaseReturnId": {NotEmpty()}}
	ReturnOrderPageInfo    = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "ReturnOrderId": {NotEmpty()}}
	Customer               = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	Authority              = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityId            = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthority           = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePassword         = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthority       = Rules{"AuthorityId": {NotEmpty()}}
)
