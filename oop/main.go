package main

import (
	"fmt"
	"go-to-hero/oop/accounts"
	"go-to-hero/oop/customers"
)

type verifyAccount interface {
	Withdraw(value float64) (string, float64, bool)
}

func PayBill(account verifyAccount, billValue float64) {
	account.Withdraw(billValue)
}

func main() {
	creatingAccounts()
}

func creatingAccounts() {
	customerJoao := customers.Customer{Name: "Joao", CPF: "00099933322", Profession: "Farmaceutico"}
	accountJoao := accounts.CheckingAccount{Holder: customerJoao, AgencyNumber: 101, AccountNumber: 998833}

	customerSilvia := customers.Customer{Name: "Silvia", CPF: "44477788811", Profession: "Vendedora"}
	accountSilvia := accounts.SavingAccount{Holder: customerSilvia, AgencyNumber: 101, AccountNumber: 998822, Operation: 1}

	accountJoao.Deposit(300)
	PayBill(&accountJoao, 100)
	fmt.Println(accountJoao.Balance())

	accountSilvia.Deposit(500)
	PayBill(&accountSilvia, 100)
	fmt.Println(accountSilvia.Balance())

}
