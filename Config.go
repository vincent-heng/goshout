package main

import (
    "encoding/json"
    "os"
    "log"
)

type Configuration struct {
    GetChatsLimit       int `json:"getChatsLimit"`
    GetMessagesLimit    int `json:"getMessagesLimit"`
}

func loadConfigurationFile() Configuration {
    file, err1 := os.Open("conf.json")
    checkErr(err1)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    checkErr(err)
    return configuration
}
