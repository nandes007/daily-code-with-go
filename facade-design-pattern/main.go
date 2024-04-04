package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacede := newWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacede.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacede.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
