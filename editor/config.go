package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

type config struct {
	DbURI        string
	StaticURI    string
	TemplatePath string
}

var conf *config

// Parse JSON config file
func parseConfig(configfile string) (*config, error) {
	data, err := ioutil.ReadFile(configfile)
	if err != nil {
		return nil, err
	}
	conf := &config{}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	dburi, err := url.Parse(strings.TrimSuffix(conf.DbURI, "/"))
	if err != nil {
		return nil, err
	}

	staticuri, err := url.Parse(strings.TrimSuffix(conf.StaticURI, "/"))
	if err != nil {
		return nil, err
	}

	if len(staticuri.String()) == 0 {
		return nil, fmt.Errorf("no static URI defined, missing staticURI entry")
	}

	templatepath := strings.TrimSuffix(conf.TemplatePath, "/")
	fi, err := os.Stat(templatepath)
	if err != nil {
		return nil, fmt.Errorf("no template path defined, missing templatePath entry: %s", err)
	}
	if fi.IsDir() == false {
		return nil, errors.New("templatePath is not a directory")
	}

	return &config{dburi.String(), staticuri.String(), templatepath}, nil
}
