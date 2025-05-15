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


```markdown
# 🛡️ VaultGuard

VaultGuard is a secure credential vault and encryption utility built to protect sensitive user data. It simulates enterprise-level **digital identity protection** and **secure data workflows**, reflecting the core security principles used in **blockchain-based payment infrastructure**.

## 🔐 Core Features

- 🔒 Encrypted vault to store user credentials
- 🔑 Password hashing and verification using SHA-256
- 🔁 Session tokens and multi-factor authorization (planned)
- 📦 Local vault with data integrity checks
- 📊 Logging and analytics for intrusion detection

## ⚙️ Tech Stack

- **Language**: Python
- **Libraries**: hashlib, uuid, os, sqlite3
- **Encryption Methods**: SHA-256, Base64 (with salt)
- **Tools**: Git, VSCode, Markdown
