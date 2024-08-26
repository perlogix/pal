package config

import (
	"errors"
	"os"
	"path/filepath"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/perlogix/pal/utils"
	"gopkg.in/yaml.v3"
)

var (
	configMap = cmap.New()
)

type Headers struct {
	Header string `yaml:"header"`
	Value  string `yaml:"value"`
}

type Upload struct {
	Enable    bool   `yaml:"enable"`
	Dir       string `yaml:"dir"`
	BasicAuth string `yaml:"basic_auth"`
}

type Config struct {
	HTTP struct {
		Listen     string `yaml:"listen"`
		TimeoutMin int    `yaml:"timeout_min"`
		Key        string `yaml:"key"`
		Cert       string `yaml:"cert"`
		Upload
	} `yaml:"http"`
	DB struct {
		EncryptKey      string    `yaml:"encrypt_key"`
		AuthHeader      string    `yaml:"auth_header"`
		ResponseHeaders []Headers `yaml:"response_headers"`
		Path            string    `yaml:"path"`
	} `yaml:"db"`
}

func InitConfig(location string) error {
	if !utils.FileExists(location) {
		return errors.New("error file does not exist: " + location)
	}

	configFile, err := os.ReadFile(filepath.Clean(location))
	if err != nil {
		return err
	}

	config := &Config{}

	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return err
	}

	configMap.Set("http_cert", config.HTTP.Cert)
	configMap.Set("http_key", config.HTTP.Key)
	configMap.Set("http_listen", config.HTTP.Listen)
	configMap.Set("http_timeout_min", config.HTTP.TimeoutMin)
	configMap.Set("http_upload", config.HTTP.Upload)
	configMap.Set("db_path", config.DB.Path)
	configMap.Set("db_encrypt_key", config.DB.EncryptKey)
	configMap.Set("db_auth_header", config.DB.AuthHeader)
	configMap.Set("db_response_headers", config.DB.ResponseHeaders)

	return nil
}

func GetConfigStr(key string) string {
	val, _ := configMap.Get(key)
	return val.(string)
}

func GetConfigInt(key string) int {
	val, _ := configMap.Get(key)
	return val.(int)
}

func GetConfigResponseHeaders() []Headers {
	val, _ := configMap.Get("db_response_headers")
	return val.([]Headers)
}

func GetConfigUpload() Upload {
	val, _ := configMap.Get("http_upload")
	return val.(Upload)
}