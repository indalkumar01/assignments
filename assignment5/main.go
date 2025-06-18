package main

import (
	"fmt"
)

type BankAccount interface {
	Init(holderName string, accountID int, accountType string, openingBalance float64)
	Withdraw(withdrawAmount float64) error
	Deposit(depositAmout float64)
	BalanceCheck()
}

type bankAccount struct {
	banlance float64
}

func NewBank() *bankAccount {
	return &bankAccount{}
}

func (b *bankAccount) Init(holderName string, accountID int, accountType string, openingBalance float64) {
	fmt.Printf("account create for %s with account id %d and account type %s with opening balance %f \n", holderName, accountID, accountType, openingBalance)
	b.banlance += openingBalance

}

func (b *bankAccount) Withdraw(withdrawAmount float64) error {
	fmt.Printf("account withdraw %.2f \n", withdrawAmount)
	b.banlance = b.banlance - withdrawAmount
	return nil
}

func (b *bankAccount) Deposit(depositAmout float64) {
	fmt.Printf("account deposit %.2f \n", depositAmout)
	b.banlance += depositAmout
}

func (b *bankAccount) BalanceCheck() {
	fmt.Println("current balanace is:=>", b.banlance)
}

func main() {
	// var i1 BankAccount
	// var i1, i2, i3, i4 BankAccount

	bank := NewBank()
	// i1 = bank

	// i1.Init("Indal", 1, "savings", 100)
	// i1.Withdraw(10)
	// i1.Deposit(100)
	// i1.BalanceCheck()

	// // 4
	// bank1 := NewBank()
	// bank2 := NewBank()
	// bank3 := NewBank()

	// i2 = bank1
	// i3 = bank2
	// i4 = bank3

	// i2.Init("Indal2", 1, "savings", 100)
	// i2.Withdraw(10)
	// i2.Deposit(200)
	// i2.BalanceCheck()

	// i3.Init("Indal2", 1, "savings", 100)
	// i3.Withdraw(10)
	// i3.Deposit(200)
	// i3.BalanceCheck()

	// i4.Init("Indal2", 1, "savings", 100)
	// i4.Withdraw(10)
	// i4.Deposit(200)
	// i4.BalanceCheck()

	// non - pointer
	newBank := NewBank2(bank)
	// newBank := NewBank2(i1)
	newBank.BalanceCheck()

}

// receiver type as parameter/polymorphism
type bankAccount2 struct {
	bankInterface BankAccount
	banlance      float64
}

func NewBank2(in BankAccount) *bankAccount2 {
	return &bankAccount2{
		bankInterface: in,
	}
}

func (b *bankAccount2) BalanceCheck() {
	b.bankInterface.BalanceCheck()
}
