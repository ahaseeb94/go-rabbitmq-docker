package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-rabbitmq-docker/src/Application/Command"
	"log"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "rabbitmq-go",
		Short: "cli to run consumers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please specify a valid command.")
		},
	}

	var startCmd = &cobra.Command{
		Use:   "fibonacci",
		Short: "Fibonacci consumer",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("start fibonacci console listener")
			listener := Command.NewFibonacciCommandService()
			listener.Execute()
		},
	}

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
