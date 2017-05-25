package main

import (
    "log"
    "net/http"
)


func main() {
	log.Printf("Starting GoShout")

	router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))

	log.Printf("Stopping GoShout")
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
