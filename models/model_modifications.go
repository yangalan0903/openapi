/*
 * JOSE Protected Message Forwarding API
 *
 * N32-f Message Forwarding Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * Source file: 3GPP TS 29.573 V16.6.0; 5G System; Public Land Mobile Network (PLMN) Interconnection; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.573/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Modifications struct {
	// Fully Qualified Domain Name
	Identity string `json:"identity" yaml:"identity" bson:"identity"`
	Operations []PatchItem `json:"operations,omitempty" yaml:"operations" bson:"operations"`
}
