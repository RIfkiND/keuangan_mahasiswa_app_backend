package main

import (
    "github.com/joho/godotenv"
    "keuangan/backend/config"
    "fmt"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    err = config.ConnectGorm()
    if err != nil {
        fmt.Println("Failed to connect to database:", err)
        return
    }
    fmt.Println("Migration successful!")
}