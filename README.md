# Bank Backend -Go

A high-performance, containerized banking backend REST API built with Go. This project simulates core banking operations such as account management, money transfers with ACID transactions, and secure authentication.

---

## 🚀 Features

* **Secure Money Transfers:** Built with database transactions (rollback support) to prevent race conditions and ensure data consistency.
* **Database:** PostgreSQL for persistent storage, with seamless migrations.
* **Containerization:** Optimized Dockerfiles using multi-stage builds for slim, secure production images.
* **Cloud-Native Deployment:** Ready for AWS EKS (Elastic Kubernetes Service) with automated secrets fetching via AWS Secrets Manager.
* **CI/CD:** Fully automated GitHub Actions pipeline for testing, linting, building, and pushing Docker images.
* **Comprehensive Testing:** * **Unit Tests:** Leverages `uber-go/mock` for mocking database interfaces.
    * **Integration Tests:** Validates real database transactions and edge cases.

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Database:** PostgreSQL
* **Migration Tools:** `golang-migrate`
* **Containerization:** Docker (Multi-stage)
* **Orchestration:** Kubernetes (AWS EKS)
* **Testing:** Go standard `testing` library + `uber-go/mock`

---

## 🚦 Getting Started

### Prerequisites

* Go (1.21+ recommended)
* Docker & Docker Compose
* `migrate` CLI (for running DB migrations)

### Local Setup

1.  **Clone the Repository**
    ```bash
    git clone [https://github.com/your-username/simple-bank.git](https://github.com/your-username/simple-bank.git)
    cd simple-bank
    ```

2.  **Spin up PostgreSQL via Docker Compose**
    ```bash
    docker-compose up -d
    ```

3.  **Run Database Migrations**
    ```bash
    # Run migrations up
    make migrateup
    ```

4.  **Run the Application**
    ```bash
    go run main.go
    ```

---

## 🧪 Testing

The project uses `uber-go/mock` to test HTTP handlers without hitting the live database.

### Run All Tests
```bash
go test -v -cover ./...
