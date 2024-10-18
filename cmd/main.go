package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"fmt"

	postgres "github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/db/postgres"
	"github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/generator/ask"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Importando o godotenv
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		ask.AskForTypeOfDB()
	},
}

var rootCmd = &cobra.Command{
	Use:   "generate-gin-schema",
	Short: "CLI API and DB Schema generator.",
	Long:  `Uma aplicação CLI simples que permite fazer perguntas e escolhas.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Por favor, use um dos comandos disponíveis.")
	},
}

var db *sql.DB

func startQuestions() {
	rootCmd.AddCommand(askCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on load file .env: %v", err)
	}

	db, err = postgres.PostgreSQLConnectDB()
	if err != nil {
		log.Fatalf("Error on connect DB (PostgreSQL): %v", err)
	}
	defer db.Close()

	if len(os.Args) > 1 {
		if os.Args[1] == "generate" {
			startQuestions()
			return
		}
	}
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API running!"))
	}).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
