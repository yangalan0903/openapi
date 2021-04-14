/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UpdatePduSessionRequest struct {
	JsonData                  *HsmfUpdateData `json:"jsonData,omitempty" multipart:"contentType:application/json"`
	BinaryDataN1SmInfoFromUe  []byte          `json:"binaryDataN1SmInfoFromUe,omitempty" multipart:"contentType:application/vnd.3gpp.5gnas,ref:JsonData.N1SmInfoFromUe.ContentId"`
	BinaryDataUnknownN1SmInfo []byte          `json:"binaryDataUnknownN1SmInfo,omitempty" multipart:"contentType:application/vnd.3gpp.5gnas,ref:JsonData.UnknownN1SmInfo.ContentId"`
}
