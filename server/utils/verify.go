package utils

var (
	IdVerify       = Rules{"ID": []string{NotEmpty()}}
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify = Rules{"PageNum": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ApiVerify      = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
)
