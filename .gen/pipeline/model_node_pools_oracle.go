/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type NodePoolsOracle struct {
	Version string `json:"version,omitempty"`
	Count int32 `json:"count,omitempty"`
	Image string `json:"image,omitempty"`
	Shape string `json:"shape,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	SubnetIds []string `json:"subnetIds,omitempty"`
}
