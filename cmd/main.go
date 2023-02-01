package main

import "github.com/erik-sostenes/accounts-api/internal/apps/backend/dependency"

func main() {
	if err := dependency.NewInjector(); err != nil {
		panic(err)
	}
}
