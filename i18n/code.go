package i18n

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// code 是唯一的
type code int

func (c code) Val() stats {
	return codeStats(c)
}

func (c code) GrpcErr() error {
	s := codeStats(c)
	return status.Error(codes.Code(s.Code), string(s.Msg))
}

func (c code) ValArgs(format string, args ...any) stats {
	s := codeStats(c)
	s.extra = fmt.Sprintf(format, args...)
	return s
}

func (c code) GrpcErrf(format string, args ...any) error {
	s := codeStats(c)
	return status.ErrorProto(s.PbState(format, args...))
}

// 通用错误
const (
	CodeUnknown       code = 10000 + iota // 未知错误
	CodeServerBusy                        // 系统繁忙
	CodeInvalidParams                     // 参数异常
	CodeNotFound                          // 资源未找到
)

const (
	CodeCaptchaUnsupportedPurpose code = 11000 + iota // 不支持的用途
	CodeCaptchaIncorrectCaptcha                       // 验证码错误
)

const (
	CodeUserUnauthorized               code = 12000 + iota // 未认证
	CodeUserInvalidToken                                   // token 无效
	CodeUserIncorrectAccountOrPassword                     // 账号或密码错误
	CodeUserInvalidPhoneNumber                             // 无效的手机号
	CodeUserNicknameOutOfRange                             // 昵称超出范围限制
	CodeUserPasswordOutOfRange                             // 密码超出范围限制
	CodeUserAlreadyRegistered                              // 账号已注册
)

const (
	CodeLedgerPriceOutOfRange code = 13000 + iota // 价格超出范围限制
	CodeLedgerTypeOutOfRange                      // 收支类型超出范围限制
)
