package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string
	Port uint16
	Tags map[string]string
}

func main() {
	c := Config{
		Host: "localhost",
		Port: 8080,
		Tags: map[string]string{"env": "dev", "test": "test1"},
	}

	bytes, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	bytes = []byte(`
host: 127.0.0.1
port: 1234
tags:
 env: dev
`)
	c = Config{}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
