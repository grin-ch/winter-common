package log_util

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	logSuffix     = ".log"
	logTimeFormat = "2006-01-02 15:04:05.000"
)

var (
	Logger = logrus.New()
)

type opt struct {
	path       string
	level      int
	jsonSwitch bool
	hasColor   bool
	hasCaller  bool
	timeFormat string
	suffix     string
	maxAge     time.Duration
}

func newOpt(opts ...Option) *opt {
	o := &opt{
		path:       "./.logs/",
		level:      int(logrus.TraceLevel),
		jsonSwitch: false,
		hasColor:   false,
		hasCaller:  true,
		timeFormat: logTimeFormat,
		suffix:     logSuffix,
		maxAge:     7 * 24 * time.Hour,
	}
	for _, fn := range opts {
		fn(o)
	}
	return o
}

type Option func(*opt)

func WithPath(path string) Option            { return func(o *opt) { o.path = path } }
func WithLevel(level int) Option             { return func(o *opt) { o.level = level } }
func WithJsonSwitch(isJson bool) Option      { return func(o *opt) { o.jsonSwitch = isJson } }
func WithColor(hasColor bool) Option         { return func(o *opt) { o.hasColor = hasColor } }
func WithCaller(hasCaller bool) Option       { return func(o *opt) { o.hasCaller = hasCaller } }
func WithTimeFormat(format string) Option    { return func(o *opt) { o.timeFormat = format } }
func WithSuffix(suffix string) Option        { return func(o *opt) { o.suffix = suffix } }
func WithMaxAge(maxAge time.Duration) Option { return func(o *opt) { o.maxAge = maxAge } }

func InitConfig(conf LogConf) {
	InitLogger(Options(conf)...)
}

// 配置日志行为
func InitLogger(opts ...Option) {
	o := newOpt(opts...)
	Logger.SetLevel(logrus.AllLevels[o.level%len(logrus.AllLevels)])
	Logger.SetReportCaller(o.hasCaller)
	if o.jsonSwitch {
		Logger.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint:      true,
			TimestampFormat:  o.timeFormat,
			CallerPrettyfier: callerPrettifier,
		})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:      o.hasColor,
			DisableColors:    !o.hasColor,
			TimestampFormat:  o.timeFormat,
			CallerPrettyfier: callerPrettifier,
		})
	}
	Logger.AddHook(fileLoggerHook(o))
}

// 文件日志
func fileLoggerHook(o *opt) logrus.Hook {
	infoWriter, _ := rotatelogs.New(o.path+"%Y%m%d.info"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(rotationTime))
	warnWriter, _ := rotatelogs.New(o.path+"%Y%m%d.warn"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(rotationTime))
	errWriter, _ := rotatelogs.New(o.path+"%Y%m%d.error"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(rotationTime))

	fileFormatter := &logrus.TextFormatter{
		ForceColors:     false,
		DisableColors:   true,
		TimestampFormat: o.timeFormat,
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errWriter,
	}, fileFormatter)
}
