package i18n

import (
	"fmt"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var modCode = map[module]map[code]stats{}

var codeMod = map[code]module{
	ok: modSystem, // 防注册
}

func register(mod module, c code, key string) {
	_, has := codeMod[c]
	if has {
		panic(fmt.Errorf("already exists code: %v, key: %v", c, key))
	}
	codeMod[c] = mod

	m, has := modCode[mod]
	if !has {
		m = map[code]stats{}
		modCode[mod] = m
	}
	m[c] = stats{
		Code: c,
		Msg:  i18nKey(fmt.Sprintf("%v.%v", mod, key)),
	}
}

var unknownStats = stats{Code: CodeUnknown, Msg: i18nUnknownErrKey}

func codeStats(c code) stats {
	mod, has := codeMod[c]
	if !has {
		return unknownStats
	}
	err, has := modCode[mod][c]
	if !has {
		return unknownStats
	}
	return err
}

type stats struct {
	Code  code    `json:"code"`
	Msg   i18nKey `json:"msg"`
	extra string
}

func (s stats) PbState(format string, args ...any) *spb.Status {
	details := make([]*anypb.Any, 0, len(args)+1)
	details = append(details, &anypb.Any{
		TypeUrl: schemeExtra,
		Value:   []byte(fmt.Sprintf(format, args...)),
	})

	return &spb.Status{
		Code:    int32(s.Code),
		Message: string(s.Msg),
		Details: details,
	}
}
func (s stats) Extra() string {
	return s.extra
}

func (s stats) String() string {
	if s.extra != "" {
		return fmt.Sprintf("code = %v msg = %v extra = %s", s.Code, s.Msg, s.extra)
	}
	return fmt.Sprintf("code = %v msg = %v", s.Code, s.Msg)
}

func IsStats(err any) (stats, bool) {
	if err == nil {
		return unknownStats, false
	}
	s, ok := err.(stats)
	return s, ok
}

const (
	schemeExtra = "extra"
)

func FromGrpcErr(err error) (stats, bool) {
	s, ok := status.FromError(err)
	if !ok {
		return unknownStats, false
	}
	extra := ""
	for _, v := range s.Proto().GetDetails() {
		if v.GetTypeUrl() == schemeExtra {
			extra = string(v.GetValue())
		}
	}
	return stats{
		Code:  code(s.Code()),
		Msg:   i18nKey(s.Message()),
		extra: extra,
	}, true
}
