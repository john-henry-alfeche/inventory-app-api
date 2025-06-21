# 📦 Inventory App API

A full-featured backend API for managing inventory operations such as products, suppliers, purchases, sales, and stock levels. Built with a focus on scalability, maintainability, and real-time stock tracking.

---

## 🚀 Features

- 🗂 Product and Category Management
- 👥 User Roles and Authentication
- 📦 Stock Movement Tracking (In/Out)
- 🛒 Purchase Order Management
- 🧾 Sales Order Processing
- 🧮 Real-Time Inventory Levels
- 🧠 Audit Logs for Transparency
- 🔒 Role-Based Access Control

---

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **Others**: UUIDs, Migrations, Environment Config

---

## 📁 Folder Structure

```plaintext
inventory-app-api/
├── cmd/                 # Entry point (main.go)
├── config/              # App configuration and env loading
├── controllers/         # HTTP handlers
├── models/              # GORM models (DB tables)
├── routes/              # API route definitions
├── services/            # Business logic
├── repositories/        # DB access and queries
├── migrations/          # Database migration scripts
├── utils/               # Helper functions, validation, etc.
├── tests/               # Unit and integration tests
└── main.go              # Application bootstrap
