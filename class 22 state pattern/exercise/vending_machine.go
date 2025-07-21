package main

import (
	"fmt"
	"log"
)

// Interface de estados
type State interface {
	InsertCoin(money int) error
	EjectCoin() (int, error)
	SelectProduct(product string) error
	Dispense() error
}

// Estados concretos
type VendingMachineContext struct {
	hasItem          State
	productRequested State
	hasMoney         State

	currentState State

	totalCoin       int
	productSelected string
	products        map[string]Product
}

type Product struct {
	name     string
	price    int
	quantity int
}

func newVendingMachine(products []Product) *VendingMachineContext {
	v := &VendingMachineContext{}
	v.products = make(map[string]Product)
	for _, p := range products {
		v.products[p.name] = p
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	productRequestedState := &ProductRequestedState{
		vendingMachine: v,
	}
	hasMoney := &HasMoneyState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.productRequested = productRequestedState
	v.hasMoney = hasMoney
	return v
}

func (v *VendingMachineContext) setState(s State) {
	v.currentState = s
}

func (v *VendingMachineContext) InsertCoin(money int) error {
	return v.currentState.InsertCoin(money)
}

func (v *VendingMachineContext) EjectCoin() (int, error) {
	return v.currentState.EjectCoin()
}

func (v *VendingMachineContext) SelectProduct(product string) error {
	return v.currentState.SelectProduct(product)
}

func (v *VendingMachineContext) Dispense() error {
	return v.currentState.Dispense()
}

// HasItemState Logica de la máquina expendedora cuando tiene productos
type HasItemState struct {
	vendingMachine *VendingMachineContext
}

func (i *HasItemState) InsertCoin(money int) error {
	if money <= 0 {
		return fmt.Errorf("Invalid coin")
	}
	i.vendingMachine.totalCoin += money
	fmt.Printf("Coin inserted: %d\n", money)
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *HasItemState) EjectCoin() (int, error) {
	if i.vendingMachine.totalCoin == 0 {
		return 0, fmt.Errorf("No coins to eject")
	}
	coin := i.vendingMachine.totalCoin
	i.vendingMachine.totalCoin = 0
	fmt.Printf("Ejected coins: %d\n", coin)
	return coin, nil
}

func (i *HasItemState) SelectProduct(product string) error {
	return fmt.Errorf("Please insert coin first")
}

func (i *HasItemState) Dispense() error {
	return fmt.Errorf("Please select product first")
}

type ProductRequestedState struct {
	vendingMachine *VendingMachineContext
}

func (i *ProductRequestedState) InsertCoin(money int) error {
	return fmt.Errorf("Product already selected, option not available")
}

func (i *ProductRequestedState) EjectCoin() (int, error) {
	if i.vendingMachine.totalCoin == 0 {
		return 0, fmt.Errorf("No coins to eject")
	}
	coin := i.vendingMachine.totalCoin
	i.vendingMachine.totalCoin = 0
	fmt.Printf("Ejected coins: %d\n", coin)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return coin, nil
}

func (i *ProductRequestedState) SelectProduct(product string) error {
	return fmt.Errorf("Product already requested")
}

func (i *ProductRequestedState) Dispense() error {
	if i.vendingMachine.totalCoin < i.vendingMachine.products[i.vendingMachine.productSelected].price {
		return fmt.Errorf("Not enough coins")
	}
	tmpProduct := i.vendingMachine.products[i.vendingMachine.productSelected]
	tmpProduct.quantity--
	i.vendingMachine.products[i.vendingMachine.productSelected] = tmpProduct
	fmt.Printf("Dispensing product: %s\n", i.vendingMachine.productSelected)
	i.vendingMachine.totalCoin -= i.vendingMachine.products[i.vendingMachine.productSelected].price
	if i.vendingMachine.totalCoin > 0 {
		fmt.Printf("Current money: %d\n", i.vendingMachine.totalCoin)
		i.vendingMachine.setState(i.vendingMachine.hasMoney)
	} else {
		fmt.Printf("Returning to initial state\n")
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	i.vendingMachine.productSelected = ""
	return nil
}

type HasMoneyState struct {
	vendingMachine *VendingMachineContext
}

func (i *HasMoneyState) InsertCoin(money int) error {
	if money <= 0 {
		return fmt.Errorf("Invalid coin")
	}
	i.vendingMachine.totalCoin += money
	fmt.Printf("Coin inserted: %d\n", money)
	return nil
}

func (i *HasMoneyState) EjectCoin() (int, error) {
	if i.vendingMachine.totalCoin == 0 {
		return 0, fmt.Errorf("No coins to eject")
	}
	coin := i.vendingMachine.totalCoin
	i.vendingMachine.totalCoin = 0
	fmt.Printf("Ejected coins: %d\n", coin)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return coin, nil
}

func (i *HasMoneyState) SelectProduct(product string) error {
	if _, ok := i.vendingMachine.products[product]; !ok {
		return fmt.Errorf("Product not available")
	}
	if i.vendingMachine.products[product].quantity == 0 {
		return fmt.Errorf("Product out of stock")
	}
	if i.vendingMachine.products[product].price > i.vendingMachine.totalCoin {
		return fmt.Errorf("Add $%v coins", i.vendingMachine.products[product].price-i.vendingMachine.totalCoin)
	}
	i.vendingMachine.productSelected = product
	fmt.Printf("Product selected: %s\n", product)
	i.vendingMachine.setState(i.vendingMachine.productRequested)
	return nil
}

func (i *HasMoneyState) Dispense() error {
	return fmt.Errorf("Please select product first")
}

func main() {
	products := []Product{
		{"Coke", 10, 5},
		{"Pepsi", 12, 3},
		{"Water", 5, 10},
	}

	vendingMachine := newVendingMachine(products)

	err := vendingMachine.InsertCoin(15)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.SelectProduct("Coke")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.Dispense()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertCoin(7)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.SelectProduct("Pepsi")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.Dispense()
	if err != nil {
		log.Fatal(err.Error())
	}

	// err = vendingMachine.InsertCoin(20)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	err = vendingMachine.SelectProduct("Cookies")
	if err != nil {
		log.Fatal(err.Error())
	}
}
