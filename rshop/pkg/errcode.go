package pkg

type ErrCode int

//@link https://github.com/darjun/errcode/blob/master

//go:generate stringer -type ErrCode -linecomment
const (
	OK                 ErrCode = iota // OK
	ERR_INVALID_PARAMS                // 无效参数
	ERR_CODE_TIMEOUT                  // 请求超时
)

func example() {
	println(OK.String())
}
