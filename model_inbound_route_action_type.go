/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// InboundRouteActionType the model 'InboundRouteActionType'
type InboundRouteActionType string

// List of InboundRouteActionType
const (
	FORWARD_TO_EMAIL InboundRouteActionType = "ForwardToEmail"
	NOTIFY_VIA_HTTP InboundRouteActionType = "NotifyViaHttp"
	STOP InboundRouteActionType = "Stop"
)

// All allowed values of InboundRouteActionType enum
var AllowedInboundRouteActionTypeEnumValues = []InboundRouteActionType{
	"ForwardToEmail",
	"NotifyViaHttp",
	"Stop",
}

func (v *InboundRouteActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InboundRouteActionType(value)
	for _, existing := range AllowedInboundRouteActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InboundRouteActionType", value)
}

// NewInboundRouteActionTypeFromValue returns a pointer to a valid InboundRouteActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInboundRouteActionTypeFromValue(v string) (*InboundRouteActionType, error) {
	ev := InboundRouteActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InboundRouteActionType: valid values are %v", v, AllowedInboundRouteActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InboundRouteActionType) IsValid() bool {
	for _, existing := range AllowedInboundRouteActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InboundRouteActionType value
func (v InboundRouteActionType) Ptr() *InboundRouteActionType {
	return &v
}

type NullableInboundRouteActionType struct {
	value *InboundRouteActionType
	isSet bool
}

func (v NullableInboundRouteActionType) Get() *InboundRouteActionType {
	return v.value
}

func (v *NullableInboundRouteActionType) Set(val *InboundRouteActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundRouteActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundRouteActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundRouteActionType(val *InboundRouteActionType) *NullableInboundRouteActionType {
	return &NullableInboundRouteActionType{value: val, isSet: true}
}

func (v NullableInboundRouteActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundRouteActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

