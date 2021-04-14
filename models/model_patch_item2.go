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

type PatchItem2 struct {
	Op PatchOperation `json:"op" yaml:"op" bson:"op"`
	Path string `json:"path" yaml:"path" bson:"path"`
	From string `json:"from,omitempty" yaml:"from" bson:"from"`
	Value map[string]interface{} `json:"value,omitempty" yaml:"value" bson:"value"`
}
