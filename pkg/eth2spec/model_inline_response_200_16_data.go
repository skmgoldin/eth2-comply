/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.11.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// InlineResponse20016Data struct for InlineResponse20016Data
type InlineResponse20016Data struct {
	// Cryptographic hash of a peer’s public key. [Read more](https://docs.libp2p.io/concepts/peer-id/)
	PeerId string `json:"peer_id,omitempty"`
	// Ethereum node record. [Read more](https://eips.ethereum.org/EIPS/eip-778)
	Enr string `json:"enr,omitempty"`
	// [Read more](https://docs.libp2p.io/reference/glossary/#multiaddr)
	Address string `json:"address,omitempty"`
	State string `json:"state,omitempty"`
	Direction string `json:"direction,omitempty"`
}
