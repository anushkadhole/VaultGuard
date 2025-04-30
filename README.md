# VaultGuard

VaultGuard is a lightweight Go CLI tool to securely store, retrieve, and delete secrets (API keys, credentials) locally, inspired by HashiCorp Vault.

## Features
- AES-256 encryption
- Role-based access (extendable)
- Time-based expiry (coming soon)

## Usage
```bash
go run main.go set
go run main.go get
go run main.go delete
```
