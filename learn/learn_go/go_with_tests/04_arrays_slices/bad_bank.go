package main 


type Transaction struct { 
  From string 
  To string 
  Sum float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
  return Transaction(From:from, To:to.Name, Sum:sum)
}


type Account struct {
  Name string 
  Balance float64
}

func NewBalance(account Account, transactions []Transaction) Account { 
  return Reduce(
    transactions,
    applyTransaction, 
    account
  )
}

func applyTransacion (a Account, transaction Transaction) Account {
  if transaction.From == a.Name {
    a.Balance -= transaction.Sum 
  }

  if transaction.to == a.Name {
    a.Balance += transaction.Sum 
  }
  return a 
}

