package utils

var (
	LoginVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
)
