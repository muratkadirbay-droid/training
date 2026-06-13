package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance >= amount {
		b.Balance -= amount
	}
}

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
}

func main() {
	a1 := &BankAccount{"Ali", 1000}
	a2 := &BankAccount{"Dias", 500}

	var acc Account

	acc = a1
	acc.Deposit(300)
	acc.Withdraw(200)

	acc = a2
	acc.Deposit(1000)
	acc.Withdraw(300)

	fmt.Printf("Счёт %s: баланс %.2f\n", a1.Owner, a1.Balance)
	fmt.Printf("Счёт %s: баланс %.2f\n", a2.Owner, a2.Balance)
}