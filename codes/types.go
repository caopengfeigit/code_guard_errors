package codes

import "google.golang.org/grpc/codes"

type ErrorCode struct {
	RpcCode    codes.Code // rpc错误码
	RpcMessage string     // rpc错误信息
	ApiCode    codes.Code // api错误码，即http code码
	ApiMessage string     // api错误信息
}

type ErrorCodes map[codes.Code]ErrorCode
