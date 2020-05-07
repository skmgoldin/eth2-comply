/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.11.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// AttesterSlashing The [`AttesterSlashing`](https://github.com/ethereum/eth2.0-specs/blob/v0.11.1/specs/core/0_beacon-chain.md#attesterslashing) object from the Eth2.0 spec.
type AttesterSlashing struct {
	Attestation1 InlineResponse20012Attestation1 `json:"attestation_1,omitempty"`
	Attestation2 InlineResponse20012Attestation1 `json:"attestation_2,omitempty"`
}
