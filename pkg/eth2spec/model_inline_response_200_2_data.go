/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.11.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// InlineResponse2002Data The [`Fork`](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/specs/phase0/beacon-chain.md#fork) object from the Eth2.0 spec.
type InlineResponse2002Data struct {
	// a fork version number
	PreviousVersion string `json:"previous_version,omitempty"`
	// a fork version number
	CurrentVersion string `json:"current_version,omitempty"`
	Epoch string `json:"epoch,omitempty"`
}
