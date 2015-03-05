package main

import (
	"encoding/json"
	"github.com/nerdalert/nfexport/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var log = logrus.New()

func setLogger(l *logrus.Logger) {
	log = l
}

func prettyPrint(data interface{}, format string) {
	var p []byte
	var err error
	switch format {
	case "json":
		p, err = json.MarshalIndent(data, "", "\t")
	case "yaml":
		p, err = yaml.Marshal(data)
	default:
		log.Printf("unsupported format: %s", format)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s", p)
}
