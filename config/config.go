package config

import (
	"fmt"
	"os"
	"time"

	"git.chinaopen.ai/yottacloud/go-libs/jwt"
	"git.chinaopen.ai/yottacloud/go-libs/redis"

	// "git.chinaopen.ai/yottacloud/tif"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

type Config struct {
	Debug    bool           `yaml:"debug,omitempty" default:"false" `
	Metrics  metricsConfig  `yaml:"metrics,omitempty"`
	Pprof    pprofConfig    `yaml:"pprof,omitempty"`
	Logger   LoggerConfig   `yaml:"logger,omitempty"`
	Database DatabaseConfig `yaml:"database,omitempty"`
	Redis    redis.Conf     `yaml:"redis,omitempty"`
	Server   ServerConfig   `yaml:"server,omitempty"`
	Jwt      JwtConfig      `yaml:"jwt,omitempty"`
	// Tif        tif.Config       `yaml:"tif,omitempty"`
	Cache      CacheConfig      `yaml:"cache,omitempty"`
	Sms        TencentSmsConfig `yaml:"sms,omitempty"`
	VerifyCode VerifyCodeConfig `yaml:"verify_code,omitempty"`
	Salt       string           `yaml:"salt,omitempty" default:"starter" `
}

type DatabaseConfig struct {
	// 仅支持 mysql
	DSN          string `yaml:"dsn"`
	MaxIdleConns int    `yaml:"max_idle_conns" default:"10"`
	MaxOpenConns int    `yaml:"max_open_conns" default:"100"`
	// format: https://golang.org/pkg/time/#ParseDuration
	ConnMaxLifetime string `yaml:"conn_max_lifetime" default:"1h"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

type ServerConfig struct {
	Host     string `yaml:"host,omitempty" default:"0.0.0.0"`
	HTTPPort string `yaml:"http_port,omitempty" default:"8080"`
	GrpcPort string `yaml:"grpc_port,omitempty" default:"8082"`
}

type CacheConfig struct {
	// 用户信息
	UserExpires string `yaml:"user_expires" default:"24h"`
	// 居信地址
	AddressExpires string `yaml:"address_expires" default:"24h"`
	// 状态信息
	StatusExpires string `yaml:"status_expires" default:""`
}

type JwtConfig struct {
	Secret   string `json:"secret,omitempty"`
	ExpireIn int    `json:"expire_in,omitempty" default:"86400"`
}

func (c CacheConfig) parseExpires(expires string) time.Duration {
	exp, err := time.ParseDuration(expires)
	if err != nil || exp.Seconds() == 0 {
		exp = time.Minute * 5
	}
	return exp
}

func (c CacheConfig) GetUserExpires() time.Duration {
	return c.parseExpires(c.UserExpires)
}

func (c CacheConfig) GetAddressExpires() time.Duration {
	return c.parseExpires(c.AddressExpires)
}

func (c CacheConfig) GetStatusExpires() time.Duration {
	return c.parseExpires(c.StatusExpires)
}

type TencentSmsConfig struct {
	AppId     int    `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	Sign      string `yaml:"sign"`
	URL       string `yaml:"url" default:"https://yun.tim.qq.com/v5/tlssmssvr/sendsms"`
}

type VerifyCodeConfig struct {
	TplId        int    `yaml:"tpl_id"`
	ExpireMinute int    `yaml:"expire_minute" default:"5"`
	CodeLength   int    `yaml:"code_length" default:"6"`
	MockCode     string `yaml:"mock_code" default:""`
	PerMinute    int    `yaml:"per_minute" default:"2"`
	PerHour      int    `yaml:"per_hour" default:"5"`
	PerDay       int    `yaml:"per_day" default:"10"`
}

type metricsConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Addr    string `yaml:"addr" default:":31999"`
}

type pprofConfig struct {
	Enabled bool   `yaml:"enabled" default:"false"`
	Addr    string `yaml:"Addr" default:"127.0.0.1:32999"`
}

func initLogger(debug bool, level, output string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zap.L().Panic("set logger level fail",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "json"

	if output != "" {
		conf.OutputPaths = []string{output}
		conf.ErrorOutputPaths = []string{output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
}

func initJwt(secret string) {
	jwt.C = jwt.Conf{
		Secret: secret,
	}
	err := jwt.Init()
	if err != nil {
		panic(fmt.Errorf("初始化 jwt 失败 %e", err))
	}
}

func initRedis(conf redis.Conf) {
	redis.C = conf
}

var C *Config

func Init(cfgFile string) {
	_ = os.Setenv("CONFIGOR_ENV_PREFIX", "-")

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	initLogger(C.Debug, C.Logger.Level, C.Logger.Output)
	initJwt(C.Jwt.Secret)
	initRedis(C.Redis)

	zap.L().Debug("loaded config")
}

func init() {
	C = &Config{}
}
