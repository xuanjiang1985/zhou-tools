package config

import (
	"embed"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Setting struct init setting.yaml
type setting struct {
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
	// Setting define new setting struct
	Setting setting
	//go:embed yaml
	fs embed.FS
)

// Setup init Setting struct
func Setup() {

	file, err := fs.Open("yaml/setting.yaml")

	if err != nil {
		log.Panicln(err)
		return
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Panicln(err)
		return
	}

	err = yaml.Unmarshal(bytes, &Setting)

	if err != nil {
		log.Panicln(err)
	}
}
