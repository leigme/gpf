package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Json struct {
	LastTemplate string `json:"last_template"`
	LastGenerate string `json:"last_generate"`
}

func (j *Json) Load() {
	data, err := os.ReadFile(Path(".config/gpf", "conf.json"))
	if err == nil {
		err = json.Unmarshal(data, j)
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func (j *Json) Update() {
	data, err := json.Marshal(j)
	if err == nil {
		if err = os.WriteFile(Path(".config/gpf", "conf.json"), data, os.ModePerm); err == nil {
			return
		}
	}
	log.Fatalln("config.Update()", err)
}

func Path(configDir, configName string) string {
	if homeDir, err := os.UserHomeDir(); err == nil {
		return filepath.Join(homeDir, configDir, configName)
	} else {
		log.Fatalln("config.Path()", err)
	}
	return ""
}
