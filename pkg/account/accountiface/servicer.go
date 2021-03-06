//

package accountiface

import (
	"github.com/Optum/dce/pkg/account"
)

// Servicer makes working with the Account Service struct easier
type Servicer interface {
	// Get returns an account from ID
	Get(ID string) (*account.Account, error)
	// Save writes the record to the dataSvc
	Save(data *account.Account) error
	// Update the Account record in DynamoDB
	Update(ID string, data *account.Account) (*account.Account, error)
	// Delete finds a given account and deletes it if it is not of status `Leased`. Returns the account.
	Delete(data *account.Account) error
	// List Get a list of accounts based on Principal ID
	List(query *account.Account) (*account.Accounts, error)
}
