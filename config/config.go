package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

// Settings is a struct yaml configuration
type Settings struct {
	Static
	DB
}

type Static struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Dir  string `yaml:"dir"`
}

type DB struct {
	Conn  string `yaml:"conn"`
	Shard string `yaml:"shard"`
	Name  string `yaml:"name"`
	Pool  int    `yaml:"pool"`
}

// Parse reads configuration file and stores values to struct variable
func Parse(cfgPath string) (cfg *Settings, err error) {
	c, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read config file %s", cfgPath)
	}

	if err = yaml.Unmarshal(c, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal config file %s", cfgPath)
	}
	log.Infof("%+v", cfg)
	return cfg, err
}

// HTTPAddr returns address for HTTP server to listen on
func (cfg *Settings) Addr() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func (db *DB) DBurl(pwd string) string {
	log.Infof(fmt.Sprintf("%s%s%s", db.Conn, pwd, db.Shard))
	return fmt.Sprintf("%s%s%s", db.Conn, pwd, db.Shard)
}
