package config

import (
	"embed"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// Setting struct for read setting.yaml
type Setting struct {
	AppName string `yaml:"appName"`
	AppEnv  string `yaml:"appEnv"`
	AppPort int    `yaml:"appPort"`

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

	fmt.Println(cfg)

	return cfg, err
}
