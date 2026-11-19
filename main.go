
package main

import (
	"fmt"
)

type Expense struct {
	Title  string
	Amount float64
}

func main() {

	var expenses []Expense
	var choice int

	for {

		fmt.Println("\n===== PERSONAL EXPENSE TRACKER =====")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View All Expenses")
		fmt.Println("3. Calculate Total Expenses")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addExpense(&expenses)

		case 2:
			viewExpenses(expenses)

		case 3:
			calculateTotal(expenses)

		case 4:
			fmt.Println("Thank you for using Expense Tracker!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}

func addExpense(expenses *[]Expense) {

	var title string
	var amount float64

	fmt.Print("Enter expense title: ")
	fmt.Scan(&title)

	fmt.Print("Enter expense amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Amount must be greater than zero.")
		return
	}

	newExpense := Expense{
		Title:  title,
		Amount: amount,
	}

	*expenses = append(*expenses, newExpense)

	fmt.Println("Expense added successfully!")
}

func viewExpenses(expenses []Expense) {

	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		return
	}

	fmt.Println("\n----- ALL EXPENSES -----")

	for i, expense := range expenses {

		fmt.Printf("%d. %s - Rs. %.2f\n",
			i+1, expense.Title, expense.Amount)
	}
}

func calculateTotal(expenses []Expense) {

	var total float64

	for _, expense := range expenses {
		total += expense.Amount
	}

	fmt.Printf("Total Expenses: Rs. %.2f\n", total)
}
