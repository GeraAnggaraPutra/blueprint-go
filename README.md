# GO Template 

## Table of Contents

- [GO Template](#go-template)
  - [Tech Stack](#tech-stack)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Local Development](#local-development)
  - [Folder Structure](#folder-structure)
 
## Tech Stack

This Go template leverages the following technologies and tools:

- **Version**: Go v1.20.
- **Framework**: Echo v4.11.
- **Database**: PostgreSQL v15.
- **Dependencies**: golang-jwt/jwt v3.2.2, google/uuid v1.6.0, jmoiron/sqlx v1.3.5, lib/pq v1.10.9

## Getting Started

To get started this go template, follow the instructions below.

### Prerequisites

- [GO](https://golang.org/dl/): Programming language.
- [Download and install PostgreSQL](https://www.postgresql.org/download/).
  - Create a database for your project and update the configuration accordingly.
- [Postman](https://www.postman.com/downloads/): Install Postman for API testing.


### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/GeraAnggaraPutra/blueprint-go
   ```

## Local Development

1. Navigate to the project directory
    ```bash
    cd your-go-template-directory
    ```

2. Copy the .env.example file to .env and update the configurations with your local settings.
    ```bash
    cp .env.example .env
    ```

3. Install the project dependencies
    ```bash
    go mod tidy
    ```

4. Run the application
    ```bash
    go run cmd/server.go
    ```

## Folder Structure

```
blueprint-go/
|-- cmd/
|-- constant/
|-- controller/
|   |-- feature/
|-- database/
|   |-- migrations/
|-- db/
|-- handler/
|   |-- auth/
|   |-- jwt/
|-- helpers/
|   |-- crypt/
|   |-- currency/
|   |-- file/
|   |-- mail/
|   |-- pdf/
|   |-- utility/
|-- middleware/
|-- module/
|-- payload/
|   |-- feature/
|-- public/
|-- repository/
|   |-- feature/
|       |-- dto/
|       |-- model/
|       |-- query/
|-- routes/
|-- service/
|   |-- feature/
|-- .env
|-- .gitignore
|-- .golangci.yaml
|-- go.mod
|-- go.sum
|-- README.md
```

<a href="https://golang.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> </a>