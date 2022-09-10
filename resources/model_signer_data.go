/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SignerData struct {
	// arbitrary stringified json object with details that will be attached to signer
	Details string `json:"details"`
	// If there are some signers with equal identity, only one signer will be chosen (either the one with the biggest weight or the one who was the first to satisfy a threshold)
	Identity uint32 `json:"identity"`
	// public key of a signer
	PublicKey string `json:"publicKey"`
	// id of the role that will be attached to a signer
	RoleId uint64 `json:"roleId"`
	// weight that signer will have, threshold for all SignerRequirements equals 1000
	Weight uint32 `json:"weight"`
}
