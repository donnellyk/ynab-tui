package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/transaction"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	c := ynab.NewClient(os.Getenv("ACCESS_TOKEN"))

	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	for _, budget := range budgets {
		fmt.Println(budget.Name + ", " + budget.ID)
		f := &transaction.Filter{
			Type: transaction.StatusUnapproved.Pointer(),
		}

		transactions, err := c.Transaction().GetTransactions(budget.ID, f)
		if err != nil {
			panic(err)
		}

		for _, transaction := range transactions {
			fmt.Println(transaction.PayeeName)
		}
	}
}
