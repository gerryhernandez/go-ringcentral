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

type CallLogRecord struct {

	// Internal identifier of a cal log record
	Id string `json:"id,omitempty"`

	// Canonical URI of a call log record
	Uri string `json:"uri,omitempty"`

	// Internal identifier of a call session
	SessionId string `json:"sessionId,omitempty"`

	// Caller information
	From CallerInfo `json:"from,omitempty"`

	// Callee information
	To CallerInfo `json:"to,omitempty"`

	// Call type
	Type_ string `json:"type,omitempty"`

	// Call direction
	Direction string `json:"direction,omitempty"`

	// Action description of the call operation
	Action string `json:"action,omitempty"`

	// Status description of the call operation
	Result string `json:"result,omitempty"`

	// The call start datetime in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	StartTime time.Time `json:"startTime,omitempty"`

	// Call duration in seconds
	Duration int32 `json:"duration,omitempty"`

	// Call recording data. Returned if the call is recorded, the withRecording parameter is set to 'True' in this case
	Recording RecordingInfo `json:"recording,omitempty"`

	// For 'Detailed' view only. The datetime when the call log record was modified in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`

	// For 'Detailed' view only. Call transport
	Transport string `json:"transport,omitempty"`

	// For 'Detailed' view only. Leg description
	Legs []LegInfo `json:"legs,omitempty"`
}