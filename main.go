package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        logrus.Info("Endpoint Hit: home")
        fmt.Fprintln(w, "Hello, World!")
    }).Methods("GET")

    logrus.Info("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

