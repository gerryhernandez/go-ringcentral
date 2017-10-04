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

type DialInNumbers struct {

	// Phone number of the dial-in number for the meeting in e.164 format
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Phone number of the dial-in number formatted according to the extension home country
	FormattedNumber string `json:"formattedNumber,omitempty"`

	// Optional field in case the dial-in number is associated with a particular location
	Location string `json:"location,omitempty"`

	// Country resource corresponding to the dial-in number
	Country DialInNumbersCountryInfo `json:"country,omitempty"`
}