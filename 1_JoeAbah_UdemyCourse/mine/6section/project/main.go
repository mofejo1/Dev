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
}

func (acc *account) GetBalance() float64 {
	return acc.Balance
}

func (acc *account) String() string {
	return fmt.Sprintf("Account Number: %s, Owner: %s, Balance: %.2f", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type savingsAccount struct {
	account
	InterestRate float64
}

func (sa *savingsAccount) ApplyInterest() {
	intrest := sa.Balance * sa.InterestRate / 100
	sa.Balance += intrest
	fmt.Printf("Adding intrest of .2%f to account %s. New Balance: %.2f\n", intrest, sa.AccountNumber, sa.Balance)

	err := sa.Deposit(intrest)
	if err != nil {
		fmt.Printf("Error applying interest to account %s: %v\n", sa.AccountNumber, err)
	}

}

type overdraftAccount struct {
	account
	OverdraftLimit float64
}

func (oa *overdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be greater than zero")
	}
	if oa.Balance+oa.OverdraftLimit < amount {
		return fmt.Errorf("Insufficient funds in account %s, including overdraft limit", oa.AccountNumber)
	}
	oa.Balance -= amount
	fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, oa.AccountNumber, oa.Balance)
	return nil
}

func main() {
	fmt.Println("Welcome to the Bank Account Management System")

	savAcc := savingsAccount{
		account: account{
			AccountNumber: "SA123456",
			Balance:       1000.00,
			OwnerName:     "Alice Smith",
		},
		InterestRate: 2.5,
	}
	fmt.Println("\n------ Savings Account Operation ")

	fmt.Println(savAcc.account.String())
}
