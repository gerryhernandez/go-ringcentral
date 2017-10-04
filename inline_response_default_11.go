/*
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type InlineResponseDefault11 struct {

	// Canonical URI of an answering rule resource
	Uri string `json:"uri,omitempty"`

	// List of answering rules
	Records []AnsweringRuleInfo `json:"records,omitempty"`

	// Information on paging
	Paging PagingInfo `json:"paging,omitempty"`

	// Information on navigation
	Navigation NavigationInfo `json:"navigation,omitempty"`
}