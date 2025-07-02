# E-COMMERCE-BASIC-GO-VANILLA

_Simplify Commerce, Empower Growth, Accelerate Success_

[![Last Commit](https://img.shields.io/github/last-commit/nawwardiab/e-commerce-basic-go-vanilla)](https://github.com/nawwardiab/e-commerce-basic-go-vanilla) [![Go](https://img.shields.io/badge/Go-82.1%25-blue)](https://golang.org/) [![Languages](https://img.shields.io/github/languages/count/nawwardiab/e-commerce-basic-go-vanilla)](https://github.com/nawwardiab/e-commerce-basic-go-vanilla)

Built with the tools and technologies:

![Markdown](https://img.shields.io/badge/Markdown-000000?logo=markdown)  
![Go Modules](https://img.shields.io/badge/Go%20Modules-000000?logo=go)  
![YAML](https://img.shields.io/badge/YAML-000000?logo=yaml)

---

## Table of Contents

- [E-COMMERCE-BASIC-GO-VANILLA](#e-commerce-basic-go-vanilla)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Testing](#testing)

---

## Overview

_e-commerce-basic-go-vanilla_ is a lightweight, modular template for building a complete e-commerce platform with vanilla Go. It provides a clean, organized architecture that simplifies development of secure user workflows, product management, and shopping-cart functionalities.

**Why _e-commerce-basic-go-vanilla_?**  
This project helps developers rapidly prototype and build scalable online stores with core features like user registration, product browsing, and cart management. The core features include:

- 🔐 **User Authentication:** Secure registration, login, and session handling to manage user identities.
- 🎨 **Dynamic Templates:** Rendered views for product listings, details, and shopping cart for a seamless user experience.
- 🚀 **Modular Architecture:** Clear separation of concerns with dedicated handlers, services, repositories, and models.
- ⚙️ **Configurable Setup:** Environment flexibility via YAML configuration for deployment across different environments.
- 🌐 **Static Asset Middleware:** Efficient static resource serving to optimize frontend performance.
- 🗄️ **Database Integration:** Reliable PostgreSQL connection management supporting persistent data storage.

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** [Go](https://golang.org/)
- **Package Manager:** Go modules

### Installation

Build _e-commerce-basic-go-vanilla_ from source and install dependencies:

1. **Clone the repository:**

```bash
git clone https://github.com/nawwardiab/e-commerce-basic-go-vanilla
```

2. **Navigate to the project directory:**

```bash
cd e-commerce-basic-go-vanilla
```

3. **Install the dependencies (using Go modules):**

```bash
go mod download
```

### Usage

Run the project with:

```bash
go run main.go
```

Or build and then run:

```bash
go build ./e-commerce-basic-go-vanilla
```

### Testing

_e-commerce-basic-go-vanilla_ uses the built-in Go testing framework. Run the full test suite with:

```bash
go test ./...
```
