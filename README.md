# Bill Buddies Server

**Bill Buddies Server** is the RESTful API backend for [Bill Buddies](https://github.com/JBenedicto13/bill-buddies-web), a playful web app that helps friends, families, and teams easily split and settle shared expenses.

---

## 🎯 Purpose

- Powers the Bill Buddies web app by managing groups, participants, expenses, and balance calculations.
- Provides API endpoints for the Next.js + TypeScript frontend.

---

## 🚀 Core Features (MVP)

- **Group Management:** Create and manage groups/events.
- **Participant Management:** Add participants to each group.
- **Expense Logging:** Record expenses (who paid, amount, description, who participated).
- **Balance Calculation:** Instantly compute who owes whom and how much.
- **Debt Settlement:** Mark debts as settled and update balances.

---

## 🏗️ Architecture & Tech Stack

- **Backend:** Go (Golang) REST API
- **Frontend:** [bill-buddies-web](https://github.com/JBenedicto13/bill-buddies-web) (Next.js + TypeScript)
- **Data Storage:** In-memory (MVP), with plans for persistent storage in future versions
- **Containerization:** Docker (recommended for development and deployment)

---

## 📐 Data Model (MVP)

```go
type Participant struct {
    Name string
}
type Expense struct {
    PaidBy       string
    Amount       float64
    Description  string
    Participants []string // names included in this expense
    Date         time.Time
}
type Group struct {
    Name         string
    Participants []Participant
    Expenses     []Expense
}
```

---

## 🚦 Getting Started

### Prerequisites
- Go: Version 1.22 or higher
- Docker: Version 27.4 or higher

### Local Development

To set up the local environment using Docker Compose, follow these steps:

1. **Ensure you have Docker and Docker Compose installed on your machine.**

2. **Clone the repo:**
   ```bash
   git clone https://github.com/JBenedicto13/bill-buddies-server.git
   cd bill-buddies-server
   ```

3. **Create a `.env` file in the root directory of your project and add the following environment variables:**

   | Variable         | Description                        |
   |------------------|------------------------------------|
   | PORT             | The port number the server will run on |
   | MODE             | The mode the server will run in (`debug` or `release`) |
   | MONGO_HOST       | The hostname for the MongoDB database |
   | MONGO_PORT       | The port number for the MongoDB database |
   | MONGO_USER       | The username for the MongoDB database |
   | MONGO_PASSWORD   | The password for the MongoDB database |
   | MONGO_DATABASE   | The name of the MongoDB database |

   Example:
   ```env
   PORT=8090
   MODE=debug
   MONGO_HOST=mongo
   MONGO_PORT=27018
   MONGO_USER=billbuddies_user
   MONGO_PASSWORD=billbuddies_password
   MONGO_DATABASE=billbuddies
   ```

4. **Run the following command to start the services:**

   ```sh
   docker compose up -d
   ```

   This command will start the following services:
   - **API Server**
   - **MongoDB**
   - **Mongo Express**

5. **Access the web app interfaces in your browser:**

   - **API Documentation**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
   - **Mongo Express**: [http://localhost:8083](http://localhost:8083)

6. **To stop the services, run:**

   ```sh
   docker compose down
   ```

---

## 📄 Documentation

The API documentation is generated using Swagger. Run the following command in the root directory of your project to regenerate the Swagger documentation:

```sh
swag init -g main.go
```

---

## 🌐 API Reference

- The API is designed for use by the [Bill Buddies Web](https://github.com/JBenedicto13/bill-buddies-web) frontend.
- [API docs coming soon!]

---

## 🤝 Contributing

- Review this overview, suggest improvements, or open issues!
- See the [frontend repo](https://github.com/JBenedicto13/bill-buddies-web) to understand how the web client consumes the API.

---

Let’s make splitting bills fun, fair, and frustration-free!
