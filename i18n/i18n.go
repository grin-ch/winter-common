package i18n

type i18nKey string

const (
	i18nUnknownErrKey i18nKey = "unknown_error"
)

type i18nKv struct {
	k code
	v string
}

func kv(code code, key string) i18nKv {
	return i18nKv{
		k: code,
		v: key,
	}
}

func init() {
	modSystem.registerKvs(
		kv(CodeUnknown, "unknown_error"),
		kv(CodeServerBusy, "server_busy"),
		kv(CodeInvalidParams, "invalid_params"),
		kv(CodeNotFound, "not_found"),
	)

	captchaUser.registerKvs(
		kv(CodeCaptchaUnsupportedPurpose, "unsupported_purpose"),
		kv(CodeCaptchaIncorrectCaptcha, "incorrect_captcha"),
	)

	modUser.registerKvs(
		kv(CodeUserUnauthorized, "unauthorized"),
		kv(CodeUserInvalidToken, "invalid_token"),
		kv(CodeUserIncorrectAccountOrPassword, "incorrect_account_or_password"),
		kv(CodeUserInvalidPhoneNumber, "invalid_phone_number"),
		kv(CodeUserNicknameOutOfRange, "nickname_out_of_range"),
		kv(CodeUserPasswordOutOfRange, "password_out_of_range"),
		kv(CodeUserAlreadyRegistered, "already_registered"),
	)

	modLedger.registerKvs(
		kv(CodeLedgerPriceOutOfRange, "price_out_of_range"),
		kv(CodeLedgerTypeOutOfRange, "type_out_of_range"),
	)
}
