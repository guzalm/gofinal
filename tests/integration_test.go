// tests/integration_test.go
package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "gofinal/tests"
)

func TestIntegrationGetProduct(t *testing.T) {
    req, err := http.NewRequest("GET", "/products/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    http.HandlerFunc(yourpackage.GetProduct).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Incorrect status code. Expected %d, got %d", http.StatusOK, status)
    }

    // Проверка ожидаемого ответа
    expected := `{"id":1,"name":"Product Name","price":100}`
    if rr.Body.String() != expected {
        t.Errorf("Incorrect response body. Expected: %s, Got: %s", expected, rr.Body.String())
    }
}
