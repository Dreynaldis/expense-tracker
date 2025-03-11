package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "expense-tracker",
		Short: "Expense tracker is a CLI tool to manage expenses",
		Long: `Manage your expenses and align it with allocated budgets for each month`,
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewBudgetCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewSummaryCmd())

	return cmd
}