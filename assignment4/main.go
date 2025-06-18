package main

import "fmt"

type BankAccount struct {
	AccountID         int
	AccountHolderName string
	AccountType       string
	OpeningBalance    float64
}

func (d *BankAccount) Initialize(accountHolderName string, accountID int, accountType string, openingBalance float64) {
	d.AccountHolderName = accountHolderName
	d.AccountID = accountID
	d.AccountType = accountType
	d.OpeningBalance = openingBalance

	fmt.Println("account initialized......")
}

func (d *BankAccount) Withdraw(num float64) {
	if d.OpeningBalance < num {
		fmt.Println("insufficient funds")
	}
	d.OpeningBalance = d.OpeningBalance - num
	fmt.Println("amount withdraw")
}

func (d *BankAccount) Deposit(num float64) {
	d.OpeningBalance += num
	fmt.Printf("deposit %.2f \n", num)
}

func (d *BankAccount) Display() {
	fmt.Println("Account ID :==>", d.AccountID)
	fmt.Println("Account Holder Name :==>", d.AccountHolderName)
	fmt.Println("opening balance:==>", d.OpeningBalance)
	fmt.Println("Account type :==>", d.AccountType)
}

func main() {
	// var accountHolderName, accountType string
	// var accountID int
	// var openingBalance float64

	// fmt.Println("Enter Account Holder Name :==> ")
	// fmt.Scan(&accountHolderName)

	// fmt.Println("Enter AccountID :==> ")
	// fmt.Scan(&accountID)

	// fmt.Println("Enter Account Type :==> ")
	// fmt.Scan(&accountType)

	// fmt.Println("Enter Balance :==> ")
	// fmt.Scan(&openingBalance)

	// fmt.Println(accountHolderName, accountID)

	// bank := BankAccount{}

	// bank.Initialize(accountHolderName, accountID, accountType, openingBalance)
	// bank.Display()

	// 3
	bankHolder := [3]BankAccount{}

	for i, h := range bankHolder {
		index := i + 1
		h.Initialize(fmt.Sprintf("indal-%d", index), index, "savings", float64(100*index))
		h.Display()
	}

}
