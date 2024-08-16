package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

   _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Initialize SQLite database
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a simple table
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    // Insert a user
    _, err = db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
    if err != nil {
        log.Fatal(err)
    }

    // Simple HTTP handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT name FROM users")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        for rows.Next() {
            var name string
            rows.Scan(&name)
            fmt.Fprintf(w, "User: %s\n", name)
        }
    })

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
