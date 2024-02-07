package main 

import (
  "errors"
  "fmt"
)

// Bicoint represents a number of Bitcoints
type Bitcoin int 


func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of bitcoint someone owns 
type Wallet struct {
  balance Bitcoin
}

//Deposit will add some Bitcoint to a waller
func (w *Wallet) Deposit(amount Bitcoin){
  w.balance += amount
}

// ErrInsufficientFunt means a wallet does not have enough Bitcoint to perform a withdraw 
var ErrInsufficientFunt = errors.New("cannot withdraw, insufficient funds")


// withdraw subtracts some Bitcoint from the wallet, returning an error if it cannot be performed
func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.balance {
    return ErrInsufficientFunt
  }

  w.balance -= ammount 
  return nil
}

// Balance returns the number of Bitcoint a wallet has. 
func (w *Wallet) Balance() Bitcoin {
  return w.balance
}





























