package i18n

type IParser interface {
	SetOptions(opts *Options)
	Parse() error
	LoadWithDefault(key string, defaultVal ...string) interface{}
	Load(keys ...string) interface{}
	LoadByLang(key, lang string) string
}
