package main

import (
    "fmt"
    "os"
    "VaultGuard/internal/crypto"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: vaultguard <command> [arguments]")
        return
    }

    command := os.Args[1]
    switch command {
    case "set":
        crypto.SetSecret()
    case "get":
        crypto.GetSecret()
    case "delete":
        crypto.DeleteSecret()
    default:
        fmt.Println("Unknown command:", command)
    }
}
