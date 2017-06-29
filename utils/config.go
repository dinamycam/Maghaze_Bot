// Package utils provides ...
package utils

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// an struct to work with the .yaml config file
type instanceConfig struct {
	botToken   string `yaml:"bot_token"`
	botDataDir string `yaml:"bot_data_dir"`
	adminPass  string `yaml:"admin_pass"`
}

func (c *instanceConfig) readconfig(conf_fname string) error {
	data, errRead := ioutil.ReadFile("./" + conf_fname)
	if errRead != nil {
		return errRead
	}

	// =| why doesn't it work?!
	errEncode := yaml.Unmarshal([]byte(data), &instanceConfig)
	if errEncode != nil {
		return errEncode
	}

	if instanceConfig.botToken == "" {
		return errors.New("config file: invalid `token`")
	} else if instanceConfig.botDataDiri == "" {
		return errors.New("config file: invalid `data_dir`")
	} else if instanceConfig.adminPass == "" {
		return errors.New("config file: invalid `admin password`")
	}
	return nil
}

func FullPath(dir string, fname string) string {
	full_path := ""
	full_dir, err := filepath.Abs(dir)
	Check(err)
	full_path = full_dir + "/" + fname
	return full_path
}
