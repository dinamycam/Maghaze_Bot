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

func (cnf *instanceConfig) readconfig(confName string) error {
	data, errRead := ioutil.ReadFile("./" + confName)
	if errRead != nil {
		return errRead
	}

	// it works, idiot me :D
	errEncode := yaml.Unmarshal([]byte(data), cnf)
	if errEncode != nil {
		return errEncode
	}

	if cnf.botToken == "" {
		return errors.New("config file: invalid `token`")
	} else if cnf.botDataDir == "" {
		return errors.New("config file: invalid `data_dir`")
	} else if cnf.adminPass == "" {
		return errors.New("config file: invalid `admin password`")
	}
	return nil
}

// FullPath takes a path and returns it's full address
func FullPath(dir string, fname string) string {
	fullPath := ""
	fullDir, err := filepath.Abs(dir)
	Check(err)
	fullPath = fullDir + "/" + fname
	return fullPath
}
