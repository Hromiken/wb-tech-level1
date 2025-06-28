package main

import "fmt"

/*21. Реализовать паттерн «адаптер» на любом примере.*/

func main() {
	Kolya := CryptoWallet{BalanceBTC: 1}
	//Gleb := EuroWallet{BalanceEUR: 322}
	adapter := &CryptoAdapter{
		Wallet: &Kolya,
		Rate:   100000,
	}

	var t Transaction = adapter
	fmt.Println("ДО")
	Kolya.Balance()
	t.Balance()
	t.Send(5436)
	fmt.Println("ПОСЛЕ")
	t.Balance()
	Kolya.Balance()
}

// EUR

type Transaction interface {
	Balance()
	Send(payment float64)
}

type EuroWallet struct {
	BalanceEUR float64
}

func (e *EuroWallet) Balance() {
	fmt.Printf("BalanceEUR: %.2f EUR\n", e.BalanceEUR)
}

func (e *EuroWallet) Send(payment float64) {
	e.BalanceEUR -= payment
	fmt.Printf("Было отправлено: %.2fEUR\nТекущий баланс: %.2f EUR\n", payment, e.BalanceEUR)
}

func (e *EuroWallet) DepositEUR(payment float64) {
	e.BalanceEUR += payment
	fmt.Printf("Поступил платеж: %.2fEUR\nТекущий баланс: %.2f EUR\n", e.BalanceEUR, e.BalanceEUR)
}

// BTC
type CryptoWallet struct {
	BalanceBTC float64
}

func (c *CryptoWallet) Balance() {
	fmt.Printf("BalanceBTC: %.2f BTC\n", c.BalanceBTC)
}

func (c *CryptoWallet) SendBTC(payment float64) {
	c.BalanceBTC -= payment
	fmt.Printf("Было отправлено: %.2f BTC\nТекущий баланс: %.2f BTC\n", payment, c.BalanceBTC)
}

func (c *CryptoWallet) DepositBTC(payment float64) {
	c.BalanceBTC += payment
	fmt.Printf("Поступил платеж: %.2f BTC\nТекущий баланс: %.2f BTC\n", payment, c.BalanceBTC)

}

// ADAPTER
type CryptoAdapter struct { // ПЕРЕВОДИТ BTC => EUR
	Wallet *CryptoWallet
	Rate   float64 // 1 BTC = 100 000 EUR (КУРС ВАЛЮТЫ)
}

func (ca *CryptoAdapter) Balance() {
	result := ca.Wallet.BalanceBTC * ca.Rate
	fmt.Printf("Balance CRYPTOWALLT in EUR = %.2f EUR\n", result)
}

func (ca *CryptoAdapter) Send(payment float64) {
	btcAmount := payment / ca.Rate
	ca.Wallet.BalanceBTC -= btcAmount
	fmt.Printf("Sent %.2f EUR (%.6f BTC)\n", payment, btcAmount)
}
