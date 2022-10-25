package utils

var (
	IdVerify          = Rules{"ID": []string{NotEmpty()}}
	LoginVerify       = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify    = Rules{"PageNum": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ApiVerify         = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	AuthorityVerify   = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify = Rules{"AuthorityId": {NotEmpty()}}
	MenuVerify        = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify    = Rules{"Title": {NotEmpty()}}
)
