package main

import "fmt"

type Bank interface {
	WithdrawMoney(accountID string, amount int) string
}

type RealBank struct{}

func (b *RealBank) WithdrawMoney(accountID string, amount int) string {
	return fmt.Sprintf("Se han retirado $%d de la cuenta %s", amount, accountID)
}

type BankProxy struct {
	realBank        *RealBank
	allowedAccounts map[string]bool
}

func NewBankProxy() *BankProxy {
	return &BankProxy{
		realBank: &RealBank{},
		allowedAccounts: map[string]bool{
			"12345": true,
			"67890": true,
		},
	}
}

func (p *BankProxy) WithdrawMoney(accountID string, amount int) string {
	if !p.allowedAccounts[accountID] {
		return fmt.Sprintf("Acceso denegado para la cuenta %s", accountID)
	}
	return p.realBank.WithdrawMoney(accountID, amount)
}

func main() {
	bankProxy := NewBankProxy()

	fmt.Println(bankProxy.WithdrawMoney("12345", 100))
	fmt.Println(bankProxy.WithdrawMoney("99999", 200))
}
