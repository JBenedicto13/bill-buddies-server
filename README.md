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

### Local Development

1. **Clone the repo:**
   ```bash
   git clone https://github.com/JBenedicto13/bill-buddies-server.git
   cd bill-buddies-server
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the server (default port: 8090):**
   ```bash
   go run main.go
   ```
   By default, the server runs on [http://localhost:8090](http://localhost:8090).

   > **Note:** You can configure the port by setting the `PORT` environment variable.

### Docker

1. **Build the Docker image:**
   ```bash
   docker build -t bill-buddies-server .
   ```

2. **Run the container (on port 8090):**
   ```bash
   docker run -p 8090:8090 --env PORT=8090 bill-buddies-server
   ```

   > **Tip:** Adjust the port if you need to avoid conflicts with other containers.

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