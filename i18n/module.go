package i18n

type module string

const (
	modSystem   module = "system"
	captchaUser module = "module_captcha"
	modUser     module = "module_user"
	modLedger   module = "module_ledger"
)

func (m module) registerKvs(kvs ...i18nKv) {
	for _, kv := range kvs {
		m.register(kv.k, kv.v)
	}
}

func (m module) register(code code, key string) {
	register(m, code, key)
}
