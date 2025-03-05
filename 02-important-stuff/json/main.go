package main

import (
	"encoding/json"
	"os"
)

type BankAccount struct {
	Number int `json:"account_number"`
	Value  int `json:"account_value"`
}

func main() {
	conta := BankAccount{Number: 1, Value: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	account2 := []byte(`{"account_number":2,"account_value":200}`)
	var contaX BankAccount
	err = json.Unmarshal(account2, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Number)
}
