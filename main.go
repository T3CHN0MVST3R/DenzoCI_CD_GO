// main.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func dialogHandler(w http.ResponseWriter, r *http.Request) {
    dialog := "What nine PLUS tEn? - Tweny one - U STPd"
    fmt.Fprintf(w, dialog)
}

func main() {
    http.HandleFunc("/", dialogHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server is running on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}
