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

type SecParamExchRspData struct {
	N32fContextId string `json:"n32fContextId" yaml:"n32fContextId" bson:"n32fContextId"`
	SelectedJweCipherSuite string `json:"selectedJweCipherSuite,omitempty" yaml:"selectedJweCipherSuite" bson:"selectedJweCipherSuite"`
	SelectedJwsCipherSuite string `json:"selectedJwsCipherSuite,omitempty" yaml:"selectedJwsCipherSuite" bson:"selectedJwsCipherSuite"`
	SelProtectionPolicyInfo *ProtectionPolicy `json:"selProtectionPolicyInfo,omitempty" yaml:"selProtectionPolicyInfo" bson:"selProtectionPolicyInfo"`
	IpxProviderSecInfoList []IpxProviderSecInfo `json:"ipxProviderSecInfoList,omitempty" yaml:"ipxProviderSecInfoList" bson:"ipxProviderSecInfoList"`
	// Fully Qualified Domain Name
	Sender string `json:"sender,omitempty" yaml:"sender" bson:"sender"`
}