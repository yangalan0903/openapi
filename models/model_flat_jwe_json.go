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

type FlatJweJson struct {
	Protected string `json:"protected,omitempty" yaml:"protected" bson:"protected"`
	Unprotected map[string]interface{} `json:"unprotected,omitempty" yaml:"unprotected" bson:"unprotected"`
	Header map[string]interface{} `json:"header,omitempty" yaml:"header" bson:"header"`
	EncryptedKey string `json:"encrypted_key,omitempty" yaml:"encrypted_key" bson:"encrypted_key"`
	Aad string `json:"aad,omitempty" yaml:"aad" bson:"aad"`
	Iv string `json:"iv,omitempty" yaml:"iv" bson:"iv"`
	Ciphertext string `json:"ciphertext" yaml:"ciphertext" bson:"ciphertext"`
	Tag string `json:"tag,omitempty" yaml:"tag" bson:"tag"`
}