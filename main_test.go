// main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestDialogHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(dialogHandler)
    handler.ServeHTTP(rr, req)

    expectedStatus := http.StatusOK
    if status := rr.Code; status != expectedStatus {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, expectedStatus)
    }

    expectedBody := "What nine PLUS tEn? - Tweny one - U STPd"
    if rr.Body.String() != expectedBody {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expectedBody)
    }
}
