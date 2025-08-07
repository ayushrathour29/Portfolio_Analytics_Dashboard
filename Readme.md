# 📊 Portfolio Analytics Dashboard

The **Portfolio Analytics Dashboard** is a full-stack web application that displays real-time statistics and performance metrics of a user's financial portfolio. It is designed to help investors monitor and evaluate their investments with clarity and ease.

---

## 🚀 Features

### Frontend Statistics Displayed

- 📈 **Portfolio Performance**: Track the overall return and absolute gain/loss of your investments.
- 🧾 **Holdings Overview**: View individual asset holdings with cost price, market value, quantity, PnL, etc.
- 🧮 **Revenue Breakdown**: Revenue calculation from each holding.
- 🧬 **Sector Allocation**: Distribution of portfolio by sectors, visualized using charts.
- 🕰️ **Historical Performance**: Timeline-based performance tracking.
- 🧠 **Insights**: Top performing asset, most held asset, underperforming sectors, etc.

---

## 🌐 RESTful API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/portfolio/holdings` | GET | Retrieves all the user's current portfolio holdings. |
| `/api/portfolio/performance` | GET | Returns key performance metrics (total investment, current value, profit/loss). |
| `/api/portfolio/summary` | GET | Provides revenue generated per holding. |
| `/api/portfolio/allocation` | GET | Returns data for sector-wise allocation of holdings. |

All endpoints return data in JSON format and are consumed by the frontend to render visualizations and tables.

---

## 🧰 Tech Stack

### Backend
- Go (Golang) with Gin Framework
- CSV data parsing
- RESTful API design

### Frontend
- React with Vite
- Tailwind CSS
- Recharts (for graphs and charts)

---

## ⚙️ How to Run the Project

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/portfolio-dashboard.git
cd portfolio-dashboard
