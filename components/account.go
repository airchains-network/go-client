package components

import (
	"log"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
)

// CheckIfAccountExists verifies if an account exists in the registry and returns its address if it does.
func CheckIfAccountExists(accountName string, client cosmosclient.Client, addressPrefix string, accountPath string) (bool, string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		log.Printf("Error creating account registry: %v", err)
		return false, ""
	}

	account, err := registry.GetByName(accountName)
	if err != nil {
		log.Printf("Account %s not found: %v", accountName, err)
		return false, ""
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Printf("Failed to get the address for account %s: %v", accountName, err)
		return false, ""
	}

	return true, addr
}

// CreateAccount creates a new account in the registry and prints its address and mnemonic.
func CreateAccount(accountName string, accountPath string) {
	registry, err := cosmosaccount.New(cosmosaccount.WithHome(accountPath))
	if err != nil {
		log.Printf("Error creating account registry: %v", err)
		return
	}

	account, mnemonic, err := registry.Create(accountName)
	if err != nil {
		log.Printf("Error creating account %s: %v", accountName, err)
		return
	}

	accountAddr, err := account.Address("air")
	if err != nil {
		log.Printf("Error retrieving address for account %s: %v", accountName, err)
		return
	}

	log.Printf("Rollup Account Created: %s", accountAddr)
	log.Printf("Mnemonic: %s", mnemonic)
}
