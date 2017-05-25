package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "html"
    "github.com/gorilla/mux"
    "io/ioutil"
    "io"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The latest version is v1")
}

func VersionIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This works, you are here: %q", html.EscapeString(r.URL.Path))
}

func ChatIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    chats := GetChats()
    err := json.NewEncoder(w).Encode(chats);
    checkErr(err)
}

func ChatShow(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    vars := mux.Vars(r)
    chatId := vars["chatId"]
    messages := GetMessages(chatId)
    err := json.NewEncoder(w).Encode(messages)
    checkErr(err)
    w.WriteHeader(http.StatusOK)
}

func ChatPost(w http.ResponseWriter, r *http.Request) {
    var message Message
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    checkErr(err)
    err2 := r.Body.Close()
    checkErr(err2)
    err3 := json.Unmarshal(body, &message)
    if err3 != nil || !isMessageValid(message) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        err4 := json.NewEncoder(w).Encode(err3)
        checkErr(err4);
    }

    vars := mux.Vars(r)
    chatId := vars["chatId"]

    PostMessage(chatId, message)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
}
