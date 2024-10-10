package main

import (
	"fmt"
    "os"
	"github.com/LeonardoGrigolettoDev/generate-gin-schema/cmd/generator/ask"
    "github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
    Use:   "meu-projeto",
    Short: "Uma CLI para perguntas e escolhas",
    Long:  `Uma aplicação CLI simples que permite fazer perguntas e escolhas.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Por favor, use um dos comandos disponíveis.")
    },
}

var askCmd = &cobra.Command{
    Use:   "action",
    Short: "Faz uma pergunta ao usuário",
    Run: func(cmd *cobra.Command, args []string) {
        ask.AskForTypeOfDB()
    },
}





func main() {
    // Adiciona o comando de perguntas
    rootCmd.AddCommand(askCmd)

    // Executa o comando raiz
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}