package main

import (
	"bytes"
	"github.com/whereabouts/sdk/config"
	"github.com/whereabouts/sdk/logger"
	"os"
	"testing"
)

var (
	json = []byte(`
	{
	 "port": 8001,
	 "name": "Korbin",
	 "age": 18
	}
	`)
	toml = []byte(`
	port = 8002
	name = "Korbin"
	age = 18
	`)
	yaml = []byte(`
	port: 8003
	name: "Korbin"
	age: 18
	`)
)

func TestLoadFile(t *testing.T) {
	file, err := os.Open("./config.json")
	if err != nil {
		logger.Fatalf("open file err: %s", err.Error())
	}
	err = config.Load(file, GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadReader(t *testing.T) {
	err := config.Load(bytes.NewBuffer(json), GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadData(t *testing.T) {
	err := config.Load(toml, GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadWithFilePath(t *testing.T) {
	configurator := config.New()
	err := configurator.Load("./config.json", GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadWithEnv(t *testing.T) {
	err := os.Setenv("CONFIG", "./config.json")
	if err != nil {
		logger.Fatalf("set env err: %s", err.Error())
	}
	err = config.LoadWithEnv("CONFIG", GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}

func TestLoadWithDefault(t *testing.T) {
	config.SetDefaultConfigPath("./config.json")
	err := config.LoadWithDefault(GetConfig())
	if err != nil {
		logger.Fatalf("load config err: %s", err.Error())
	}
	logger.Infof("%+v", *GetConfig())
}