/*
 * Elastic Email REST API
 *
 * This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    To start using this API, you will need your Access Token (available <a href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a href=\"https://api.elasticemail.com/public/help\">here</a>.
 *
 * API version: 4.0.0
 * Contact: support@elasticemail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// SplitOptimizationType the model 'SplitOptimizationType'
type SplitOptimizationType string

// List of SplitOptimizationType
const (
	OPENS SplitOptimizationType = "Opens"
	CLICKS SplitOptimizationType = "Clicks"
)

func (v *SplitOptimizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SplitOptimizationType(value)
	for _, existing := range []SplitOptimizationType{ "Opens", "Clicks",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SplitOptimizationType", value)
}

// Ptr returns reference to SplitOptimizationType value
func (v SplitOptimizationType) Ptr() *SplitOptimizationType {
	return &v
}

type NullableSplitOptimizationType struct {
	value *SplitOptimizationType
	isSet bool
}

func (v NullableSplitOptimizationType) Get() *SplitOptimizationType {
	return v.value
}

func (v *NullableSplitOptimizationType) Set(val *SplitOptimizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitOptimizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitOptimizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitOptimizationType(val *SplitOptimizationType) *NullableSplitOptimizationType {
	return &NullableSplitOptimizationType{value: val, isSet: true}
}

func (v NullableSplitOptimizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitOptimizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
