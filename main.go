// main.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func dialogHandler(w http.ResponseWriter, r *http.Request) {
    dialog := "What nine PLUS tEn? - Tweny one - U STPd"
    fmt.Fprintf(w, dialog)
}

func main() {
    http.HandleFunc("/", dialogHandler)
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
