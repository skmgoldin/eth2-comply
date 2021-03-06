/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.11.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// InlineResponse20012Attestation1 The [`IndexedAttestation`](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/specs/phase0/beacon-chain.md#indexedattestation) object from the Eth2.0 spec.
type InlineResponse20012Attestation1 struct {
	// Attesting validator indices
	AttestingIndices []string `json:"attesting_indices,omitempty"`
	Signature string `json:"signature,omitempty"`
	Data InlineResponse20011Data `json:"data,omitempty"`
}
