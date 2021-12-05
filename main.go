package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.bmvs.io/ynab"
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
	}
}
