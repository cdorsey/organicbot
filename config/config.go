package config

import (
    "github.com/cdorsey/organicbot/utils"
    "encoding/json"
    "io/ioutil"
)

type Config struct {
    Channel     string `json:"channel"`
    Nick        string `json:"nick"`
    TwitchKey  string `json:"twitch_key"`
    Debug       bool   `json:"debug"`
}

var instance *Config

func Getconfig() *Config {
    if instance == nil {
        raw, err := ioutil.ReadFile("./config/config.json"); utils.HandleError(err, true)

        json.Unmarshal(raw, &instance)
    }

    return instance
}