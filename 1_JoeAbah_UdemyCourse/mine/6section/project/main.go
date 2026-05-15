package main

import "fmt"

type account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Deposit amount must be greater than zero")
	}
	acc.Balance += amount
	fmt.Printf("Deposited %.2f to account %s. New balance: %.2f\n", amount, acc.AccountNumber, acc.Balance)
	return nil
}


func (acc *account) Withdraw(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("Withdrawal amount must be greater than zero")
    }
    if acc.Balance < amount {
        return fmt.Errorf("Insufficient funds in account %s", acc.AccountNumber)
    }
    acc.Balance -= amount
    fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, acc.AccountNumber, acc.Balance)
    return nil
}func main() {

}
