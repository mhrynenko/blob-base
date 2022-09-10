/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateAccountOp struct {
	// ID of account to be created
	Destination string `json:"destination"`
	// ID of an another account that introduced this account into the system. If account with such ID does not exist or it's Admin Account. Referrer won't be set.
	Referrer string `json:"referrer"`
	// ID of the role that will be attached to an account
	RoleId uint64 `json:"roleId"`
	// Array of data about 'destination' account signers to be created
	SignersData []SignerData `json:"signersData"`
}
