// main_test.go
package main

import (
    "example.com/gofinal/tests"
)

func TestGenerateOTP(t *testing.T) {
    otp := GenerateOTP()

    // Проверка длины сгенерированного OTP
    if len(otp) != 6 {
        t.Errorf("Incorrect OTP length. Expected: 6, Got: %d", len(otp))
    }

    // Проверка, что OTP состоит только из цифр
    _, err := strconv.Atoi(otp)
    if err != nil {
        t.Errorf("Invalid OTP format. Expected: all digits, Got: %s", otp)
    }
}


func TestIntegrationGetProduct(t *testing.T) {
    tests.TestIntegrationGetProduct(t)
}