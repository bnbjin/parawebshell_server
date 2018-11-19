package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	dftConfigDir  = "/tmp/paraws/"
	dftConfigName = "config.json"
)

type Config struct {
	reserved string
}

var (
	cfginst Config
)

func init() {
	err := cfginst.read(dftConfigDir)
	if nil != err {
		panic(err)
	}
}

func (c *Config) read(path string) error {
	log.Println("start reading configuration ", path)
	defer func() {
		log.Println("finish reading configuration")
	}()

	data, err := ioutil.ReadFile(path)
	if nil != err {
		log.Fatal("reading configuration ", path, " with ", err)
		return err
	}

	err = json.Unmarshal(data, &cfginst)
	if nil != err {
		return err
	}

	return nil
}
