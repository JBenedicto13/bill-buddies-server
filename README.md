# Bill Buddies

**Bill Buddies** is a playful, web-based split bill tracker designed to help friends, families, and teams easily manage shared expenses and settle up with confidence.

---

## 🎯 Purpose

- Make group expense tracking simple, transparent, and fair
- Calculate who owes whom and by how much
- Help friends avoid awkward money conversations after shared events

---

## 🚀 Core Features (MVP)

1. **Create Groups:**  
   Start a new group for an event (e.g., trip, dinner, shared apartment).

2. **Add Participants:**  
   Add people by name to each group.

3. **Log Expenses:**  
   - Record what was spent, who paid, and who participated in each expense.
   - Expenses are split evenly in the MVP (custom splits in future versions).

4. **View Balances:**  
   Instantly see who owes whom and how much.

5. **Settle Debts:**  
   Mark debts as paid and update group balances.

---

## 🤓 Example User Flow

1. Alice creates a "Beach Trip" group with Bob and Cara.
2. Alice pays $90 for groceries (all participate).
3. Bob pays $30 for gas (all participate).
4. Cara pays $60 for boat rental (all participate).
5. App calculates and displays:
   - Who owes whom and by how much, summarized simply.

---

## 🛠️ Tech Stack (Proposed)

- **Backend:** Go (Golang) – REST API
- **Frontend:** (To be decided: initial version could be Go templates, React, Svelte, or Vue)
- **Data Storage:** In-memory for MVP, with plans for file-based or lightweight DB
- **Version Control:** GitHub

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

## 📏 Next Steps

- Finalize MVP feature set
- Sketch main UI screens/user flows
- Set up project repos for backend and frontend
- Break down MVP into tasks/issues

---

## 🤝 Contributing

- Review this overview and suggest improvements!
- We’ll divide core features for implementation after finalizing requirements.

---

Let’s make splitting bills fun, fair, and frustration-free!