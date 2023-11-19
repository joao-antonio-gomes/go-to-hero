package accounts

import (
	"go-to-hero/oop/customers"
	"strconv"
)

type CheckingAccount struct {
	Holder                      customers.Customer
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CheckingAccount) Withdraw(value float64) (string, float64, bool) {
	haveBalance := c.balance >= value
	isValidValue := value > 0

	if haveBalance && isValidValue {
		c.balance -= value
		return "Saque realizado com sucesso.", c.balance, true
	}

	return "Saldo insuficiente.", 0, false
}

func (c *CheckingAccount) Deposit(value float64) (string, float64, bool) {
	isValidValue := value > 0
	if isValidValue {
		c.balance += value
		return "Valor depositado com sucesso.", c.balance, true
	}

	return "Valor inválido.", 0, false
}

func (c *CheckingAccount) Transfer(value float64, targetAccount *CheckingAccount) (string, float64, float64) {
	withdrawMessage, _, withdrawSuccess := c.Withdraw(value)
	if !withdrawSuccess {
		return withdrawMessage, c.balance, targetAccount.balance
	}

	depositMessage, _, depositSuccess := targetAccount.Deposit(value)
	if !depositSuccess {
		return depositMessage, c.balance, targetAccount.balance
	}

	return "Transferência realizada com sucesso.", c.balance, targetAccount.balance
}

func (c *CheckingAccount) Balance() string {
	return "Seu saldo é de R$ " + strconv.FormatFloat(c.balance, 'f', 2, 64)
}
