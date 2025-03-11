package cmd

import (
	"github.com/dreynaldis/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var BudgetMonth int32
var BudgetAmount float64

func NewBudgetCmd() *cobra.Command{
	budgetCmd := &cobra.Command{
		Use: "budget",
		Short: "Budget for the month",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunBudgetCmd(args)
		},
	}
	budgetCmd.Flags().Int32VarP(&BudgetMonth, "month", "m", 0, "Month for the expense summary")
	budgetCmd.MarkFlagRequired("month")
	budgetCmd.Flags().Float64VarP(&BudgetAmount, "amount", "a", 0, "Budget for the month")
	budgetCmd.MarkFlagRequired("amount")

	return budgetCmd
}

func RunBudgetCmd(args []string) error {
	return expense.BudgetMonth(BudgetMonth, BudgetAmount)
}