/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type VsmfUpdateError struct {
	Error                 *ProblemDetails  `json:"error"`
	Pti                   int32            `json:"pti,omitempty"`
	N1smCause             string           `json:"n1smCause,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData `json:"unknownN1SmInfo,omitempty"`
	FailedToAssignEbiList []int32          `json:"failedToAssignEbiList,omitempty"`
	NgApCause             *NgApCause       `json:"ngApCause,omitempty"`
	Var5gMmCauseValue     int32            `json:"5gMmCauseValue,omitempty"`
	RecoveryTime          *time.Time       `json:"recoveryTime,omitempty"`
}
