package expense

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/dreynaldis/expense-tracker/internal/log"
)

type Expense struct {
	ID          int64 `json:"id"`
	Description string `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
}

func NewExpense (id int64, description string, amount float64, category string) *Expense {
	return &Expense{
		ID: id,
		Description: description,
		Amount: amount,
		Category: category,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func AddExpense(description string, amount float64, category string) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	var newExpenseId int64
	if len(expenses) == 0 {
		newExpenseId = 1
	} else {
		newExpenseId = expenses[len(expenses)-1].ID + 1
	}

	thisMonth := time.Now().Month()
	thisMonthBudget, err := GetMonthlyBudget(int32(thisMonth))
	if err != nil{
		return err
	}

	newExpense := NewExpense(newExpenseId,description, amount, category )
	expenses = append(expenses, *newExpense)

	thisMonthExpenses := 0.0
	for _, expense := range expenses {
		if expense.CreatedAt.Month() == thisMonth {
			thisMonthExpenses += expense.Amount
		}
	}

	if thisMonthBudget != 0 && thisMonthExpenses > thisMonthBudget {
		log.Warning(fmt.Sprintf("You have exceeded this month's budget. Budget : %.2f, Expenses : %.2f", thisMonthBudget, thisMonthExpenses))
	}

	log.Info(fmt.Sprintf("Expense added: %s, Amount: %.2f, Category: %s (ID: %d)", description, amount, category, newExpenseId))
	return WriteExpensesToFile(expenses)
}

func ListExpenses(category string) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		log.Error("No Expenses found")
		return nil
	}

	var filteredExpenses []Expense
	for _, expense := range expenses {
		if category == "all" || strings.EqualFold(expense.Category, category) {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}
	if len(filteredExpenses) == 0 {
		log.Info(fmt.Sprintf("No Expenses found for category : %s", category))
		return nil
	}

	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			Margin(1,0).
			Render(fmt.Sprintf("Expenses: (%s)", category)))
	for _, expense := range filteredExpenses {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Width(8).
			Render(fmt.Sprintf("ID: %d", expense.ID))
		
		expensesStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("%s %s %.2f %s", formattedId, expense.Description, expense.Amount, expense.Category))
		
		fmt.Println(expensesStyle)
	}

	return nil
}

func SummaryExpenses(month int) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		log.Error("No expenses found")
		return nil
	}

	var total float64
	if month == 0 {
		for _, expense := range expenses {
			total += expense.Amount
		}
	} else {
		for _, expense := range expenses {
			if expense.CreatedAt.Month() == time.Month(month) {
				total += expense.Amount
			}
		}
	}
	log.Info(fmt.Sprintf("Total expenses : %.2f", total))
	return nil
}

func DeleteExpense(id int64) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	var found bool = false
	for i, expense := range expenses {
		if expense.ID == id {
			found = true
			expenses = append(expenses[:i], expenses[i+1:]...)
			log.Info(fmt.Sprintf("Expenses deleted: %s, Amount: %.2f, Category: %s, (ID: %d)", expense.Description, expense.Amount, expense.Category, expense.ID))
			break
		}
	}

	if !found {
		log.Error(fmt.Sprintf("Expense not found with ID: %d", id))
		return nil
		
	}
	return WriteExpensesToFile(expenses)
}