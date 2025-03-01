package i18n

const ok code = 200 // 正常

type rsp struct {
	Code code
	Data any `json:",omitempty"`
}

func Success(data any) rsp {
	return rsp{
		Code: ok,
		Data: data,
	}
}

func NilSuccess() rsp {
	return Success(nil)
}
