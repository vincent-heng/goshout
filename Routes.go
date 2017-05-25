package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Index",
        "GET",
        "/v1",
        VersionIndex,
    },
    Route{
        "ChatIndex",
        "GET",
        "/v1/chats",
        ChatIndex,
    },
    Route{
        "ChatShow",
        "GET",
        "/v1/chats/{chatId}",
        ChatShow,
    },
    Route{
        "ChatPost",
        "POST",
        "/v1/chats/{chatId}",
        ChatPost,
    },
}
