package cmd

import (
	"github.com/dreynaldis/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var listCategory string

func NewListCmd() *cobra.Command {
	listCmd:= &cobra.Command{
		Use: "list",
		Short: "List all expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListExpenseCmd(args)
		},
	}
	listCmd.Flags().StringVarP(&listCategory, "category", "c", "all", "Filter Expense by Category")
	return listCmd
}


func RunListExpenseCmd(args []string) error {
	return expense.ListExpenses(listCategory)
}