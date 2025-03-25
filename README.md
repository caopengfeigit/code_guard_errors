# 云码卫士错误码库
## 遵循规则
```go
type ErrorCode struct {
    RpcCode codes.Code // rpc错误码
    RpcMessage string // rpc错误信息
    ApiCode codes.Code // api错误码，即http code码
    ApiMessage string // api错误信息
}
```
## code码规则
```
[A]（1 位）：错误级别
1：客户端错误（4xx）
2：服务器错误（5xx）
3：第三方服务错误
9：未知错误

[BBB]（3 位）：服务模块编号（不同微服务有唯一编号）
001：系统服务（rpc-sys）
999：通用错误

[CCC]（3 位）：具体错误类型<br>
001：资源未找到（NotFound）<br>
002：无权限（PermissionDenied）<br>
003：参数错误（InvalidArgument）<br>
500：服务器内部错误（InternalError）
```