package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	fileName = pflag.StringP("file", "f", "", "The toml file must be parse")
)

type Config struct {
	Title    string
	Owner    OwnerConfig
	Database DatabaseConfig
	Servers  map[string]ServerConfig
	Clients  ClientConfig
}

type OwnerConfig struct {
	Name string
	DOB  time.Time
}

type DatabaseConfig struct {
	Server        string
	Ports         []int
	ConnectionMax int `toml:"connection_max"`
	Enabled       bool
}

type ServerConfig struct {
	IP string
	DC string
}

type ClientConfig struct {
	Data  [][]interface{}
	Hosts []string
}

func main() {
	pflag.Parse()

	if *fileName == "" {
		logrus.Fatal("filename is empty")
		os.Exit(1)
	}

	var config Config

	if _, err := toml.DecodeFile(*fileName, &config); err != nil {
		fmt.Println(err)
		return
	}

	spew.Dump(config)
}
