/*
 * N32 Handshake API
 *
 * N32-c Handshake Service.  © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  All rights reserved. 
 *
 * Source file: 3GPP TS 29.573 V16.5.0; 5G System; Public Land Mobile Network (PLMN) Interconnection; Stage 3
 * Url: http://www.3gpp.org/ftp/Specs/archive/29_series/29.573/
 *
 * API version: 1.1.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ProtectionPolicy struct {
	ApiIeMappingList []ApiIeMapping `json:"apiIeMappingList" yaml:"apiIeMappingList" bson:"apiIeMappingList"`
	DataTypeEncPolicy []IeType `json:"dataTypeEncPolicy,omitempty" yaml:"dataTypeEncPolicy" bson:"dataTypeEncPolicy"`
}
