package options

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Version string = "v0.0.1"

type LogOptions struct {
	Level    string `yaml:"level"`
	Path     string `yaml:"path"`
	MaxSize  int    `yaml:"maxsize"`
	MaxAge   int    `yaml:"maxage"`
	Compress bool   `yaml:"compress"`
}

type CoreOptions struct {
	Mode string     `yaml:"mode"`
	Log  LogOptions `yaml:"log"`
}

type MetricsOptions struct {
	Enable bool   `yaml:"enable"`
	Path   string `yaml:"path"`
}

type RedisOptions struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	DialTimeout int    `yaml:"dialtimeout"`
}

type APIServerOptions struct {
	Listen  string         `yaml:"listen"`
	Redis   RedisOptions   `yaml:"redis"`
	Metrics MetricsOptions `yaml:"metrics"`
	Prefix  string         `yaml:"prefix"`
}

type Options struct {
	Core      CoreOptions      `yaml:"core"`
	APIServer APIServerOptions `yaml:"apiserver"`
}

func NewOptions() (opts Options) {
	optsSource := viper.AllSettings()
	err := createOptions(optsSource, &opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create options failed:", err)
		os.Exit(1)
	}
	return
}
