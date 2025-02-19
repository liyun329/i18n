package i18n

import (
	"fmt"

	"sync"
)

// Options 配置
type Options struct {
	DefaultParser   string
	DefaultLang     string
	LangDirectory   string
	CacheDirectory  string
	EnableFileAsKey bool
}

// Option 配置驱动
type Option func(*Options)

type I18n struct {
	opts   *Options
	parser *Parser
}

var once sync.Once
var i18nInit *I18n

// NewI18n 初始化包对象函数
func NewI18n(opt ...Option) *I18n {
	once.Do(func() {
		i18nInit = &I18n{
			opts: &Options{},
		}
	})
	// 初始化配置
	i18nInit.initOption(opt...)

	// 初始化解析器
	err := i18nInit.initParser()
	if err != nil {
		panic(err)
	}

	return i18nInit
}

func (i *I18n) initOption(opt ...Option) {
	for _, o := range opt {
		o(i.opts)
	}
}

func (i *I18n) initParser() error {
	// 检查是否设置了解析器, 如果没有, 则默认使用json解析器
	if i.opts.DefaultParser == "" {
		i.initOption(DefaultParser("json"))
	}
	// 检查是否设置了语言, 如果没有, 则默认使用 zh_cn
	if i.opts.DefaultLang == "" {
		i.initOption(DefaultLang("zh_cn"))
	}
	// 加载解析器
	i.parser = NewParser()

	var parser = i.parser.Getter(i.opts.DefaultParser)
	if parser == nil {
		return fmt.Errorf("未注册解析器")
	}

	// 传入配置
	parser.SetOptions(i.opts)
	// 解析内容
	err := parser.Parse()
	return err
}

func (i *I18n) Load(keys ...string) interface{} {
	var parser = i.parser.Getter(i.opts.DefaultParser)
	if parser == nil {
		panic(fmt.Errorf("未注册的解析器"))
	}

	//// 传入配置
	//parser.SetOptions(i.opts)
	//// 解析内容
	//err := parser.Parse()
	//return err
	return parser.Load(keys...)
}

func (i *I18n) LoadByLang(key, lang string) string {
	var parser = i.parser.Getter(i.opts.DefaultParser)
	if parser == nil {
		panic(fmt.Errorf("未注册的解析器"))
	}

	//// 传入配置
	//parser.SetOptions(i.opts)
	//// 解析内容
	//err := parser.Parse()
	//return err
	return parser.LoadByLang(key, lang)
}

func (i *I18n) LoadWithDefault(key string, defaultVal ...string) interface{} {
	var parser = i.parser.Getter(i.opts.DefaultParser)
	if parser == nil {
		panic(fmt.Errorf("未注册的解析器"))
	}

	//// 传入配置
	//parser.SetOptions(i.opts)
	//// 解析内容
	//err := parser.Parse()
	//return err
	return parser.LoadWithDefault(key, defaultVal...)
}

// Lang 修改语言
func (i *I18n) Lang(l string) {
	i.opts.DefaultLang = l
}

// LangDirectory 存放不同语言的目录
func LangDirectory(ld string) Option {
	return func(o *Options) {
		o.LangDirectory = ld
	}
}

// DefaultLang 默认语言
func DefaultLang(l string) Option {
	return func(o *Options) {
		o.DefaultLang = l
	}
}

// DefaultParser 默认配置解析器
func DefaultParser(p string) Option {
	return func(o *Options) {
		o.DefaultParser = p
	}
}

// CacheDirectory 配置解析器后的缓存目录
func CacheDirectory(p string) Option {
	return func(o *Options) {
		o.DefaultParser = p
	}
}

// EnableFileAsKey 将文件名作为其中的key
func EnableFileAsKey(p bool) Option {
	return func(o *Options) {
		o.EnableFileAsKey = p
	}
}
