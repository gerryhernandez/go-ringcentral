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

import (
	"time"
)

type RangesInfo struct {

	// Date and time in format YYYY-MM-DD hh:mm
	From time.Time `json:"from,omitempty"`

	// Date and time in format YYYY-MM-DD hh:mm
	To time.Time `json:"to,omitempty"`
}