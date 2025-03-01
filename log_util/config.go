package log_util

import "time"

const rotationTime = 24 * time.Hour // 按天分割

type LogConf struct {
	Level      int  `json:",optional"`
	HasColor   bool `json:",optional"`
	HasCaller  bool `json:",optional"`
	JsonSwitch bool `json:",optional"`

	Path       string `json:",optional"`
	Suffix     string `json:",optional"`
	TimeFormat string `json:",optional"`
	MaxDays    int    `json:",optional"`
}

func Options(conf LogConf) []Option {
	opts := []Option{
		WithLevel(conf.Level),
		WithColor(conf.HasColor),
		WithCaller(conf.HasCaller),
		WithJsonSwitch(conf.JsonSwitch),
	}
	if conf.Path != "" {
		opts = append(opts, WithPath(conf.Path))
	}
	if conf.Suffix != "" {
		opts = append(opts, WithSuffix(conf.Suffix))
	}
	if conf.TimeFormat != "" {
		opts = append(opts, WithTimeFormat(conf.TimeFormat))
	}
	if conf.MaxDays > 0 {
		opts = append(opts, WithMaxAge(time.Duration(conf.MaxDays)*rotationTime))
	}
	return opts
}
