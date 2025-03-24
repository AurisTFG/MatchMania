# 🏎 MatchMania

**A competitive matchmaking app for the racing game Trackmania**, originally developed as a university project for **Web Application Design (T120B165)** class and later extended as a **Bachelor’s Degree Final Project (P000B001)**

## 🎯 Purpose

**MatchMania** aims to deliver a dynamic and balanced matchmaking experience for Trackmania players. By implementing an **ELO-based ranking system**, it ensures fair competition by matching players of similar skill levels, enhancing engagement and competitive integrity.

## 🚀 Functional Requirements

- **ELO Ranking System**  
  Tracks player performance and adjusts rankings based on match outcomes, ensuring competitive balance through continuous recalibration.

- **User Registration & Profiles**  
  Players can create an account, log in, and manage secure profiles displaying their statistics, match history, and overall rankings.

- **Season Management**  
  Supports the creation of multiple competitive seasons, each with its own ELO-based ranking system for a fresh start and renewed competition.

- **Team Registration & Matchmaking**  
  Allows players to form teams and participate in team-based or individual matchmaking, ensuring balanced pairings based on skill level.

## 🛠️ Technologies Used

### Backend

- **Go** – High-performance programming language.
  - **Gin** – Lightweight HTTP web framework.
  - **GORM** – ORM for database management.
- **JWT** – Secure authentication using JSON Web Tokens.

### Frontend

- **pnpm** – Fast and efficient package manager.
- **Vite** – Modern build tool for quick development and optimized production.
- **React** – Frontend framework.
  - **TypeScript** – Type-safe development.
  - **SWC** – Fast JSX/TypeScript compilation.

### General

- **PostgreSQL** – High-performance relational database for storing user data, match results, and rankings.
- **Makefile** – Automates build, run, test, clean and other tasks.
- **Docker** – Containerizes the app for consistent and portable environments.

## 📄 Documentation

The OpenAPI Swagger specification is available at:  
[backend/docs/swagger.yaml](../backend/docs/swagger.yaml)

## ▶️ How to Run

1. **Clone the repository**:

   ```sh
   git clone https://github.com/auristfg/matchmania.git
   ```

2. **Run Backend**:

   - Navigate to the `backend` directory:

     ```sh
     cd .\backend\
     ```

   - Run the API using Makefile:

     ```sh
     make run-dev
     ```

3. **Run Frontend**:

   - Navigate to the `frontend` directory:

     ```sh
     cd .\frontend\
     ```

   - Run the web app using Makefile:

     ```sh
     make run-dev
     ```
