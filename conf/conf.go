package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/labstack/gommon/log"
)

var (
	Conf              config // holds the global app config.
	defaultConfigFile = "conf/config.toml"
)

// config 는 설정
type config struct {
	ReleaseMode bool   `toml:"release_mode"`
	LogLevel    string `toml:"log_level"`

	SessionStore string `toml:"session_store"`
	CacheStore   string `toml:"cache_store"`

	App app

	Server   server
	SMTP     smtp     `toml:"smtp"`
	Database database `toml:"database"`
}

type app struct {
	Name      string `toml:"name"`
	JwtSecret string `toml:"jwt_secret"`
}

type aws struct {
	S3BucketName string `toml:"s3_bucket_name"`
	S3AccessKey  string `toml:"s3_access_key"`
	S3SecretKey  string `toml:"s3_secret_key"`
}

type server struct {
	Graceful bool   `toml:"graceful"`
	Protocol string `toml:"protocol"`
	Addr     string `toml:"addr"`

	Domain       string `toml:"domain"`
	DomainAPI    string `toml:"domain_api"`
	DomainWeb    string `toml:"domain_web"`
	DomainSocket string `toml:"domain_socket"`
}

type static struct {
	Type string `toml:"type"`
}

type smtp struct {
	Server string `toml:"server"`
	Port   string `toml:"port"`
	From   string `toml:"from"`
	Pwd    string `toml:"pwd"`
}

type database struct {
	Adapter   string `toml:"adapter"`
	Name      string `toml:"name"`
	UserName  string `toml:"user_name"`
	Pwd       string `toml:"pwd"`
	Host      string `toml:"host"`
	Port      int    `toml:"port"`
	IdleConns int    `toml:"idle_conns"`
	OpenConns int    `toml:"open_conns"`
	LogMode   bool   `toml:"log_mode"`
}

func init() {
}

// InitConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	if configFile == "" {
		configFile = defaultConfigFile
	}

	// Set defaults.
	Conf = config{
		ReleaseMode: false,
	}

	if _, err := os.Stat(configFile); err != nil {
		return errors.New("config file err:" + err.Error())
	}

	// log.Infof("load config from file:" + configFile)
	configBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}

	return nil
}

func GetLogLvl() log.Lvl {
	//DEBUG INFO WARN ERROR OFF
	switch Conf.LogLevel {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OF":
		return log.OFF
	}

	return log.DEBUG
}
