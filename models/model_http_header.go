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

type HttpHeader struct {
	Header string `json:"header" yaml:"header" bson:"header"`
	Value *EncodedHttpHeaderValue `json:"value" yaml:"value" bson:"value"`
}
