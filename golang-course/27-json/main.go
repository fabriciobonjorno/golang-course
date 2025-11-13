package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"` // json tags to specify field names in JSON
	Balance int `json:"b"`
}

func main() {
	acc := Account{Number: 12345, Balance: 1000}
	res, err := json.Marshal(acc) // Convert Account struct to JSON wtith Marshal you need a variable
	if err != nil {
		panic(err)
	}
	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(acc); err != nil { // Encode and write JSON to stdout directly with Encode
		panic(err)
	}

	var acc2 Account
	if err := json.Unmarshal(res, &acc2); err != nil { // Convert JSON back to Account struct with Unmarshal you need a pointer
		panic(err)
	}
	println(acc2.Number, acc2.Balance)

	var acc3 Account
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&acc3); err != nil { // Read JSON from stdin and convert to Account struct with Decode you need a pointer
		panic(err)
	}
	println(acc3.Number, acc3.Balance)
}
