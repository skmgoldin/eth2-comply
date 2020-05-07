/*
 * Eth2 Beacon Node API
 *
 * API specification for the beacon node, which enables users to query and participate in Ethereum 2.0 phase 0 beacon chain.
 *
 * API version: Dev - Eth2Spec v0.11.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eth2spec
// InlineResponse2003Data struct for InlineResponse2003Data
type InlineResponse2003Data struct {
	PreviousJustified InlineResponse2003DataPreviousJustified `json:"previous_justified,omitempty"`
	CurrentJustified InlineResponse2003DataPreviousJustified `json:"current_justified,omitempty"`
	Finalized InlineResponse2003DataPreviousJustified `json:"finalized,omitempty"`
}
