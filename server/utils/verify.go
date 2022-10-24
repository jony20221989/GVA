package utils

var (
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify = Rules{"PageNum": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
