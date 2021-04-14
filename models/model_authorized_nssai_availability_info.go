/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthorizedNssaiAvailabilityInfo struct {
	AuthorizedNssaiAvailabilityData []AuthorizedNssaiAvailabilityData `json:"authorizedNssaiAvailabilityData" bson:"authorizedNssaiAvailabilityData"`

	SupportedFeatures string `json:"supportedFeatures,omitempty" bson:"supportedFeatures"`
}