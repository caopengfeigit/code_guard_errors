package codes

import "net/http"

var Codes = ErrorCodes{
	0: {
		RpcCode:    0,
		RpcMessage: "success",
		ApiCode:    http.StatusOK,
		ApiMessage: "success",
	},
}
