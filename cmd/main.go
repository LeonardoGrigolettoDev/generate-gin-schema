package main

import (
    "log"
    "net/http"
    "database/sql"
	postgres "github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/db/postgres"
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"  // Importando o godotenv
)

var db *sql.DB


func main() {
    // Carregando variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error on load file .env: %v", err)
    }

    db, err = postgres.PostgreSQLConnectDB()
    if err != nil {
        log.Fatalf("Error on connect DB (PostgreSQL): %v", err)
    }
    defer db.Close()

    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("API running!"))
    }).Methods("GET")

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
