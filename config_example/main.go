package main

import (
	"github.com/getevo/evo"
	"log"
)

type Custom struct {
	Array []string `yaml:"array"`
}

func main() {
	evo.Setup()
	var cfg Custom

	//load custom key from custom file
	evo.ParseConfig("custom.yml", "custom", &cfg)
	log.Print(cfg)

	cfg = Custom{}
	//load custom key from app config
	evo.ParseConfig("", "custom", &cfg)
	log.Print(cfg)

}

///
