package app

import (
	"flag"
	"log"

	"github.com/larspensjo/config"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

var TOPIC = make(map[string]string)

func LoadConfig() {
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("CONF") {
		section, err := cfg.SectionOptions("CONF")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("CONF", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	return
}
