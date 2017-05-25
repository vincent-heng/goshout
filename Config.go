package main

import (
    "encoding/json"
    "os"
)

type Configuration struct {
    getChatsLimit       int
    getMessagesLimit    int
}

func loadConfigurationFile() Configuration { // FIXME returns {0,0}
    file, err1 := os.Open("conf.json")
    checkErr(err1)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    checkErr(err)
    return configuration
}
