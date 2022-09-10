/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type TransactionEnvelope struct {
	Key
	Attributes TransactionEnvelopeAttributes `json:"attributes"`
}
type TransactionEnvelopeResponse struct {
	Data     TransactionEnvelope `json:"data"`
	Included Included            `json:"included"`
}

type TransactionEnvelopeListResponse struct {
	Data     []TransactionEnvelope `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustTransactionEnvelope - returns TransactionEnvelope from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTransactionEnvelope(key Key) *TransactionEnvelope {
	var transactionEnvelope TransactionEnvelope
	if c.tryFindEntry(key, &transactionEnvelope) {
		return &transactionEnvelope
	}
	return nil
}
