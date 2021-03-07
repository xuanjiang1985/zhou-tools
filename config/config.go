package config

import (
	"embed"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Setting struct for read setting.yaml
type Setting struct {
	AppName string `yaml:"appName"`
	AppEnv  string `yaml:"appEnv"`
	AppPort int    `yaml:"appPort"`

	Log struct {
		Level  string
		Logdir string
	}

	Database struct {
		Host   string
		User   string
		Dbname string
		Pwd    string
	}

	Redis struct {
		Host string
		Pwd  string
	}
}

var (
	//go:embed yaml
	fs embed.FS
)

// Read returns yaml setting data
func Read() (cfg Setting, err error) {

	file, err := fs.Open("yaml/setting.yaml")

	if err != nil {
		return cfg, errors.WithStack(err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return cfg, errors.WithStack(err)
	}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		return cfg, errors.WithStack(err)
	}

	return cfg, err
}
