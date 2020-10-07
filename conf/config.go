package conf

import (
	"encoding/json"
	"io/ioutil"
)

type Db struct {
	Path string `json:"Path"`
}

func (d *Db) GetDb() string {
	return d.Path
}

type config struct {
	Db *Db `json:"db"`
}

var cfg *config

func InitConfig(path string) error {
	var (
		b   []byte
		err error
	)
	cfg = new(config)
	b, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, cfg)
}

func GetConfig() *config {
	return cfg
}

func (c *config) GetDb() string {
	return c.Db.Path
}
