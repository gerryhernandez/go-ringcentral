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

type InlineResponseDefault6 struct {

	// List of countries which can be selected for a dialing plan
	Records []DialingPlanCountryInfo `json:"records,omitempty"`

	// Information on paging
	Paging PagingInfo `json:"paging,omitempty"`

	// Information on navigation
	Navigation NavigationInfo `json:"navigation,omitempty"`
}