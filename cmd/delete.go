package cmd

import (
	"github.com/dreynaldis/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var deleteExpenseId int64

func NewDeleteCmd() *cobra.Command{
	deleteCmd := &cobra.Command{
		Use: "delete",
		Short: "Delete Expenses by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteExpensesCmd(args)
		},
	}
	
	deleteCmd.Flags().Int64VarP(&deleteExpenseId, "id", "i", 0, "ID of the expense to delete")
	return deleteCmd
}

func RunDeleteExpensesCmd(args []string) error {
	return expense.DeleteExpense(deleteExpenseId)
}