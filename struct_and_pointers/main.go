package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
	fmt.Println("Successfully Deposited, current balance is:", a.Balance)
}

func (a *Account) Withdraw(amount int) {
	if amount > a.Balance {
		fmt.Println("Insufucient Balance")
	} else {
		a.Balance -= amount
		fmt.Println("Successfully Withdrwan Balance, current balance is:", a.Balance)
	}
}

func main() {
	pialAccount := Account{Owner: "Pial", Balance: 0}

	pialAccount.Deposit(2000)

	fmt.Println(pialAccount.Balance)

	pialAccount.Withdraw(4000)
	fmt.Println(pialAccount.Balance)
}
